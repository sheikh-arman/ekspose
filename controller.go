type controller struct{
	clientset kubernetes.Interface
	depLister appslisters.DeploymentLister
	depCacheSynced cache.InformerSynced
	queue workqueue.RateLimitingInterface

}
func newController(clientset kubernetes.Interface, depInformer appsinformers.DeploymentInformer) *controller{
	c:= &controller{
		cache.ResourceEventHandlerFuncs{
			AddFunc: c.handleAdd,
			DeleteFunc: c.handleDel,
		},
	}
	return c
}