# Port Service

Stores and provides information about maritime ports around the world.

For each port we have the following data:
- UN/LOCODE
- Name
- City
- Country
- any aliases
- coordinates
- province
- timezone
- code if applicable

# Contributing
Run `scripts/onboarding.sh` to set up go grpc protoc

Whenever you change any proto files run `scripts/generate-stubs.sh` to regenerate the `.pb.go` files