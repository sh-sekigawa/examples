package dnspopulate

import (
	"context"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connection"
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/networkservice"
	"github.com/networkservicemesh/networkservicemesh/sdk/endpoint"

	"github.com/sirupsen/logrus"
	
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"github.com/sh-sekigawa/dnsendpoint/api/types/v1alpha1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

const (
	DNSEndpointNamespace = "nsm-dns"
)

type DNSEndpointInfo struct {
	name string
	namespace string
	hostName string
	recordType string
	ipAddress string
}

type DNSEndpointPair struct {
	NSEEndpoint *DNSEndpointInfo
	NSCEndpoint *DNSEndpointInfo
}

// CustomFuncEndpoint is endpoint that apply passed ConnectionMutator to connection that accepts from next endpoint
type DNSPopulateEndpoint struct {
	name string
	client *rest.RESTClient
	endpoint string
	searchDomain string
	dnsEndpoints map[string]*DNSEndpointPair
}

func (info *DNSEndpointInfo) Deploy(client *rest.RESTClient) (v1alpha1.DNSEndpoint, error) {
	dnsEndpoint := v1alpha1.DNSEndpoint {
		ObjectMeta: metav1.ObjectMeta{
			Name: info.name,
		},
		Spec: v1alpha1.DNSEndpointSpec {
			Endpoints: []*v1alpha1.Endpoint {
				v1alpha1.NewEndpoint(info.hostName, info.recordType, info.ipAddress),
			},
		},
	}

	result := v1alpha1.DNSEndpoint{}
	err := client.Post().
		Namespace(info.namespace).Resource("dnsendpoints").
		Body(&dnsEndpoint).
		Do().Into(&result)
	return result, err
}

func (info *DNSEndpointInfo) Delete(client *rest.RESTClient) (error) {
	err := client.Delete().
			Namespace(info.namespace).Resource("dnsendpoints").
			Name(info.name).
			Body(&metav1.DeleteOptions{}).
			Do().Error()
	return err
}

// Request implements Request method from NetworkServiceServer
// Consumes from ctx context.Context:
//	   Next
func (cf *DNSPopulateEndpoint) Request(ctx context.Context, request *networkservice.NetworkServiceRequest) (*connection.Connection, error) {
	Log(ctx).Infof("DNS populater request on connection: %v", request.GetConnection())
	id := request.Connection.Id
	nseAppName := cf.endpoint
	nscAppName := request.Connection.Labels["app"]
	networkservicename := request.Connection.NetworkService
	dnsEndpointPairName := nscAppName + "-" + nseAppName + "-" + networkservicename + "-" + id
	dnsEndpointPair := DNSEndpointPair{
		NSEEndpoint: &DNSEndpointInfo{
			name: nseAppName + "-" + networkservicename + "-" + id,
			namespace: DNSEndpointNamespace,
			hostName: nseAppName + "." + networkservicename + "." + cf.searchDomain,
			recordType: v1alpha1.RecordTypeA,
			ipAddress: strings.Split(request.Connection.Context.IpContext.DstIpAddr, "/")[0],
		},
		NSCEndpoint: &DNSEndpointInfo{
			name: nscAppName + "-" + networkservicename + "-" + id,
			namespace: DNSEndpointNamespace,
			hostName: nscAppName + "." + networkservicename + "." + cf.searchDomain,
			recordType: v1alpha1.RecordTypeA,
			ipAddress: strings.Split(request.Connection.Context.IpContext.SrcIpAddr, "/")[0],
		},
	}

	// Create NSE DNSEndpoint
	result, err := dnsEndpointPair.NSEEndpoint.Deploy(cf.client)
	if err != nil {
		Log(ctx).Error(err.Error())
		return nil, err
	} else {
		Log(ctx).Infof("%v\n", result)
	}
	// Create NSC DNSEndpoint
	result, err = dnsEndpointPair.NSCEndpoint.Deploy(cf.client)
	if err != nil {
		Log(ctx).Error(err.Error())
		// Delete NSE DNSEndpoint if error occurs.
		errDel := dnsEndpointPair.NSEEndpoint.Delete(cf.client)
		if errDel != nil {
			Log(ctx).Error(errDel.Error())
		}
		return nil, err
	} else {
		Log(ctx).Infof("%v\n", result)
		// Add DNSEndpointPair to DNSPopulateEndpoint
		cf.dnsEndpoints[dnsEndpointPairName] = &dnsEndpointPair
	}

	// resultList := v1alpha1.DNSEndpointList{}
	// err = client.Get().
	// 	Namespace("nsm-dns").Resource("dnsendpoints").
	// 	Do().Into(&resultList)
	// if err != nil {
	// 	Log(ctx).Error(err.Error())
	// }else{
	// 	Log(ctx).Infof("%v\n", resultList)
	// }
	
	// pods, err := client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	// if err != nil {
	// 	Log(ctx).Error(err.Error())
	// }
	// Log(ctx).Infof("There are %d pods in the cluster\n", len(pods.Items))

	// // Examples for error handling:
	// // - Use helper functions e.g. errors.IsNotFound()
	// // - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	// _, err = client.CoreV1().Pods("default").Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})
	// if errors.IsNotFound(err) {
	// 	Log(ctx).Info("Pod example-xxxxx not found in default namespace\n")
	// } else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	// 	Log(ctx).Infof("Error getting pod %v\n", statusError.ErrStatus.Message)
	// } else if err != nil {
	// 	Log(ctx).Error(err.Error())
	// } else {
	// 	Log(ctx).Info("Found example-xxxxx pod in default namespace\n")
	// }
	
	if endpoint.Next(ctx) != nil {
		return endpoint.Next(ctx).Request(ctx, request)
	}

	return request.GetConnection(), nil
}

// Close implements Close method from NetworkServiceServer
// Consumes from ctx context.Context:
//	   Next
func (cf *DNSPopulateEndpoint) Close(ctx context.Context, connection *connection.Connection) (*empty.Empty, error) {
	Log(ctx).Infof("DNS populater completed on connection: %v", connection) 
	id := connection.Id
	nseAppName := cf.endpoint
	nscAppName := connection.Labels["app"]
	networkservicename := connection.NetworkService
	dnsEndpointPairName := nscAppName + "-" + nseAppName + "-" + networkservicename + "-" + id
	if cf.dnsEndpoints[dnsEndpointPairName] != nil {
		if cf.dnsEndpoints[dnsEndpointPairName].NSCEndpoint != nil {
			err := cf.dnsEndpoints[dnsEndpointPairName].NSCEndpoint.Delete(cf.client)
			if err != nil {
				Log(ctx).Error(err.Error())
			}
		}
		if cf.dnsEndpoints[dnsEndpointPairName].NSEEndpoint != nil {
			err := cf.dnsEndpoints[dnsEndpointPairName].NSEEndpoint.Delete(cf.client)
			if err != nil {
				Log(ctx).Error(err.Error())
			}
		}
		Log(ctx).Infof("DNS Endpoint Pair %s deleted\n", dnsEndpointPairName)
		// Delete DNSEndpointPair from DNSPopulateEndpoint
		delete(cf.dnsEndpoints, dnsEndpointPairName)
	} else {
		// Already deleted.
		Log(ctx).Errorf("DNS Endpoint Pair %s is already deleted¥n", dnsEndpointPairName)
	}
	
	if endpoint.Next(ctx) != nil {
		return endpoint.Next(ctx).Close(ctx, connection)
	}
	return &empty.Empty{}, nil
}

// Name returns the composite name
func (cf *DNSPopulateEndpoint) Name() string {
	return "dnspopulater"
}

func NewDNSPopulateEndpoint(name string, client *rest.RESTClient, endpoint string, searchDomain string, dnsEndpoints *(map[string]*DNSEndpointPair)) *DNSPopulateEndpoint {
	return &DNSPopulateEndpoint{
		name: name,
		client: client,
		endpoint: endpoint,
		searchDomain: searchDomain,
		dnsEndpoints: *dnsEndpoints,
	}
}

func NewKubernetesClient() (*rest.RESTClient, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	
	v1alpha1.AddToScheme(scheme.Scheme)

    crdConfig := *config
    crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: v1alpha1.GroupName, Version: v1alpha1.GroupVersion}
    crdConfig.APIPath = "/apis"
    crdConfig.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)
    crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

    return rest.UnversionedRESTClientFor(&crdConfig)
}

func Log(ctx context.Context) logrus.FieldLogger {
	if rv, ok := ctx.Value("Log").(logrus.FieldLogger); ok {
		return rv
	}
	return logrus.New()
}