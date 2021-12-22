# TODO: Replce this with helm dependencies

helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo add jetstack https://charts.jetstack.io

helm upgrade ingress-nginx ingress-nginx/ingress-nginx \
  --install \
  --create-namespace \
  --namespace="ingress-nginx"

helm upgrade cert-manager jetstack/cert-manager \
  --install \
  --namespace cert-manager \
  --create-namespace \
  --set serviceAccount.annotations."eks\.amazonaws\.com/role-arn"=$(terraform output cert_manager_irsa_role_arn) \
  --set installCRDs=true \
  --set securityContext.enabled=true \
  --set securityContext.fsGroup=1001 \
  --wait