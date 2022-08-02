package main

import (
	"context"
	"fmt"
	"testing"

	// client "github.com/kubernetes-sdk-for-go-101/pkg/client"
	kosmosv1 "github.com/kosmos-industrie40/setting-up-kosmos-edge/api/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	fake "sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestCreateObject(t *testing.T) {

	var testCases = []struct {
		name            string
		id              int
		spec            kosmosv1.KosmosKubeEdgeSpec
		pods            []runtime.Object
		expectedSuccess bool
	}{
		{
			name: "all_needed_values_given",
			id:   123,
			spec: kosmosv1.KosmosKubeEdgeSpec{
				Body: kosmosv1.BodyProperties{
					Machine: "machine-1",
				},
				Signature: kosmosv1.SignatureProperties{
					Meta: kosmosv1.MetaSignature{
						Algorithm: "algorithm",
						Date:      "today",
					},
					Signature: "signature",
				},
			},
			pods: []runtime.Object{
				&corev1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "kosmos-contract-test",
						Namespace: "default",
					},
				},
			},
			expectedSuccess: true,
		},
	}

	client := fake.NewClientBuilder().WithScheme(scheme).Build()

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			err := createObject(client, test.id, test.spec)
			if err != nil {
				t.Fatalf("unexpected error creating object: %v", err)
			}

			contract := &kosmosv1.KosmosKubeEdge{}
			err = client.Get(context.Background(), types.NamespacedName{
				Namespace: "default",
				Name:      fmt.Sprintf("kosmos-contract-%d", test.id),
			}, contract)
			if err != nil {
				t.Fatalf("unexpected error creating object: %v", err)
			}

			// Test values are set
			if contract.ObjectMeta.Namespace != "default" {
				t.Fatalf("Expected default was :%v", contract.ObjectMeta.Namespace)
			}
			if contract.ObjectMeta.Name != fmt.Sprintf("kosmos-contract-%d", test.id) {
				t.Fatalf("Expected kosmos-contract-%d was :%v", test.id, contract.ObjectMeta.Name)
			}
			if contract.TypeMeta.Kind != "KosmosKubeEdge" {
				t.Fatalf("Expected KosmosKubeEdge but was: %v", contract.TypeMeta.Kind)
			}
			if contract.Spec.Body.Machine != "machine-1" {
				t.Fatalf("Expected maschine-1 but was: %v", contract.Spec.Body.Machine)
			}

		})
	}

}
