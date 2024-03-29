helm upgrade butler ../helm/services \
  --install \
  --set image.tag=$RELEASE_VERSION \
  --set image.repository=$ECR_REPO \
  --set certManager.route53.hostedZoneID=$HOSTED_ZONE_ID \
  --set services.users.postgres.host=$USERS_DB_HOST \
  --set services.users.postgres.name=$USERS_DB_NAME \
  --set services.users.postgres.user=$USERS_DB_USER \
  --set services.users.postgres.password=$USERS_DB_PASSWORD \
  --set services.users.jwtSecret=$USERS_JWT_SECRET \
  --set environment=prod \
  --wait