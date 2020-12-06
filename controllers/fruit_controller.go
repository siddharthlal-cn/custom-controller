/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"fmt"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	plantsv1alpha1 "siddharth.com/plant-factory/api/v1alpha1"
)

// FruitReconciler reconciles a Fruit object
type FruitReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=plants.siddharth.com,resources=fruits,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=plants.siddharth.com,resources=fruits/status,verbs=get;update;patch

func (r *FruitReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("fruit", req.NamespacedName)

	var fruit plantsv1alpha1.Fruit
	if err := r.Get(ctx, req.NamespacedName, &fruit); err != nil {
		log.Error(err, "Failed to get fruit")

		return ctrl.Result{}, err
	}

	desiredFruit := createDesiredFruitDeployment(fruit)
	if err := r.Create(ctx, desiredFruit); err != nil {
		log.Error(err, "Failed to create fruit deployment")
		return ctrl.Result{}, err
	}

	fruit.Status.Created = true
	if err := r.Status().Update(ctx, &fruit); err != nil {
		log.Error(err, "Failed to update fruit status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *FruitReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&plantsv1alpha1.Fruit{}).
		Complete(r)
}

func createDesiredFruitDeployment(fruit plantsv1alpha1.Fruit) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fruit.ObjectMeta.Name,
			Namespace: fruit.ObjectMeta.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &fruit.Spec.Amount,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"plantsv1alpha1/fruit": fruit.ObjectMeta.Name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fruit.ObjectMeta.Name,
					Namespace: fruit.ObjectMeta.Namespace,
					Labels: map[string]string{
						"plantsv1alpha1/fruit": fruit.ObjectMeta.Name,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "fruit-container",
							Image: "nginx:1.14.2",
							Env: []corev1.EnvVar{
								{
									Name:  "HELLO_MSG",
									Value: fmt.Sprintf("from a %s", fruit.Spec.Type),
								},
							},
						},
					},
				},
			},
		},
	}
}
