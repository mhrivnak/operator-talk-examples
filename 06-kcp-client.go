import (
    "context"
    kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
    ctrl "sigs.k8s.io/controller-runtime"
)

func (r *reconciler)  Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    ctx = kcpclient.WithCluster(ctx, req.ObjectKey.Cluster)
	// from here on pass this context to all client calls
    //		...
}
