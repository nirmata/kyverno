package context

import (
	"reflect"
	"testing"

	urkyverno "github.com/kyverno/kyverno/api/kyverno/v1beta1"
	"github.com/kyverno/kyverno/pkg/config"
	"github.com/kyverno/kyverno/pkg/engine/jmespath"
	kubeutils "github.com/kyverno/kyverno/pkg/utils/kube"
	"github.com/stretchr/testify/assert"
	authenticationv1 "k8s.io/api/authentication/v1"
)

var (
	jp  = jmespath.New(config.NewDefaultConfiguration(false))
	cfg = config.NewDefaultConfiguration(false)
)

func Test_addResourceAndUserContext(t *testing.T) {
	var err error
	rawResource := []byte(`
	{
		"apiVersion": "v1",
		"kind": "Pod",
		"metadata": {
		   "name": "image-with-hostpath",
		   "labels": {
			  "app.type": "prod",
			  "namespace": "my-namespace"
		   }
		},
		"spec": {
		   "containers": [
			  {
				 "name": "image-with-hostpath",
				 "image": "docker.io/nautiker/curl",
				 "volumeMounts": [
					{
					   "name": "var-lib-etcd",
					   "mountPath": "/var/lib"
					}
				 ]
			  }
		   ],
		   "volumes": [
			  {
				 "name": "var-lib-etcd",
				 "emptyDir": {}
			  }
		   ]
		}
	 }
			`)

	userInfo := authenticationv1.UserInfo{
		Username: "system:serviceaccount:nirmata:user1",
		UID:      "014fbff9a07c",
	}
	userRequestInfo := urkyverno.RequestInfo{
		Roles:             nil,
		ClusterRoles:      nil,
		AdmissionUserInfo: userInfo,
	}

	var expectedResult string
	ctx := NewContext(jp)
	err = AddResource(ctx, rawResource)
	if err != nil {
		t.Error(err)
	}
	result, err := ctx.Query("request.object.apiVersion")
	if err != nil {
		t.Error(err)
	}
	expectedResult = "v1"
	t.Log(result)
	if !reflect.DeepEqual(expectedResult, result) {
		t.Error("exected result does not match")
	}

	err = ctx.AddUserInfo(userRequestInfo)
	if err != nil {
		t.Error(err)
	}
	result, err = ctx.Query("request.object.apiVersion")
	if err != nil {
		t.Error(err)
	}
	expectedResult = "v1"
	t.Log(result)
	if !reflect.DeepEqual(expectedResult, result) {
		t.Error("exected result does not match")
	}

	result, err = ctx.Query("request.userInfo.username")
	if err != nil {
		t.Error(err)
	}
	expectedResult = "system:serviceaccount:nirmata:user1"
	t.Log(result)
	if !reflect.DeepEqual(expectedResult, result) {
		t.Error("exected result does not match")
	}
	// Add service account Name
	err = ctx.AddServiceAccount(userRequestInfo.AdmissionUserInfo.Username)
	if err != nil {
		t.Error(err)
	}
	result, err = ctx.Query("serviceAccountName")
	if err != nil {
		t.Error(err)
	}
	expectedResult = "user1"
	t.Log(result)
	if !reflect.DeepEqual(expectedResult, result) {
		t.Error("exected result does not match")
	}

	// Add service account Namespace
	result, err = ctx.Query("serviceAccountNamespace")
	if err != nil {
		t.Error(err)
	}
	expectedResult = "nirmata"
	t.Log(result)
	if !reflect.DeepEqual(expectedResult, result) {
		t.Error("expected result does not match")
	}
}

func Test_ImageInfoLoader(t *testing.T) {
	resource1, err := kubeutils.BytesToUnstructured([]byte(`{
		"apiVersion": "v1",
		"kind": "Pod",
		"metadata": {
		  "name": "test-pod",
		  "namespace": "default"
		},
		"spec": {
		  "containers": [{
			"name": "test_container",
			"image": "nginx:latest"
		  }]
		}
	}`))
	assert.Nil(t, err)
	newctx := newContext()
	err = newctx.AddImageInfos(resource1, cfg)
	assert.Nil(t, err)
	// images not loaded
	assert.Nil(t, newctx.images)
	// images loaded on Query
	name, err := newctx.Query("images.containers.test_container.name")
	assert.Nil(t, err)
	assert.Equal(t, name, "nginx")
}

func Test_ImageInfoLoader_OnDirectCall(t *testing.T) {
	resource1, err := kubeutils.BytesToUnstructured([]byte(`{
		"apiVersion": "v1",
		"kind": "Pod",
		"metadata": {
		  "name": "test-pod",
		  "namespace": "default"
		},
		"spec": {
		  "containers": [{
			"name": "test_container",
			"image": "nginx:latest"
		  }]
		}
	}`))
	assert.Nil(t, err)
	newctx := newContext()
	err = newctx.AddImageInfos(resource1, cfg)
	assert.Nil(t, err)
	// images not loaded
	assert.Nil(t, newctx.images)
	// images loaded on explicit call to ImageInfo
	imageinfos := newctx.ImageInfo()
	assert.Equal(t, imageinfos["containers"]["test_container"].Name, "nginx")
}
