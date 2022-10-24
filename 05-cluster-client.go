func (r *MyReconciler) addCluster(name string, kubeconfig []byte) {
    config, _ := clientcmd.RESTConfigFromKubeConfig(kubeconfig)
    newCluster, _ := cluster.New(config)
    // this can be done even after the Manager starts
    mgr.Add(newCluster)
    // must keep your own reference to this cluster for later use
    r.clusters[name] = newCluster
    r.controller.Watch(
        source.NewKindWithCache(&corev1.ConfigMap{}, newCluster.GetCache()),
        // probably need a custom handler to map the remote cluster's event to
        // a primary resource in the main cluster.
        &myhandlers.CustomHandler{},
    )
}

func (r *MyReconciler) getRemoteConfigMap(clusterName string, key types.NamespacedName) *corev1.ConfigMap {
    // find the Cluster wherever you stashed it and get its Client
    client := r.cluster[name].GetClient()
    cm := corev1.ConfigMap{}
    client.Get(&cm, key)
    return &cm
}
