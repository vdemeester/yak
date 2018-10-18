# yak — Yet Another Kubernetes …

My hack-a-day tool to manage kubernetes while I'm developping. My use
case is the following : I'm developping on and on top of Kubernetes,
thus I tend to create/destroy/… kubernetes a lot.

This is kind-of a *pot-pourri* of commands

## TODOs

- [ ] Provision a kubernetes cluster (`yak m/mk/mo`)
  - [ ] `minikube`
  - [ ] `minishift`
  - [ ] … ?`
- [ ] `kube-prompt`-like
- [ ] Configuration management
  - [ ] *hook* into minikube/minishift (how they add stuff, etc)
  - [ ] `kubectx`/`kubeens`-like
- [ ] `kubectl` command-line management (aka get same version as
	  server if possible)
- [ ] `plumini` ideas (watch, diff, …)
- [ ] integration with other tools (`skaffold`, `ko`, `kustomize`,
	  `odo`, …) — shell out ?
  - [x] if `cmd` is in lookup path (and there is no)
  - [ ] integrate with tmux [ex](https://raw.githubusercontent.com/awslabs/aws-service-operator/master/aws-service-operator-example.gif).
