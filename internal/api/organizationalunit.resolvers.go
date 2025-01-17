package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"

	"entgo.io/contrib/entgql"
	"go.infratographer.com/resource-provider-api/internal/ent/generated"
	"go.infratographer.com/resource-provider-api/internal/ent/generated/resourceprovider"
	"go.infratographer.com/x/gidx"
)

// ResourceProvider is the resolver for the resourceProvider field.
func (r *organizationalUnitResolver) ResourceProvider(ctx context.Context, obj *OrganizationalUnit, after *entgql.Cursor[gidx.PrefixedID], first *int, before *entgql.Cursor[gidx.PrefixedID], last *int, orderBy *generated.ResourceProviderOrder, where *generated.ResourceProviderWhereInput) (*generated.ResourceProviderConnection, error) {
	return r.client.ResourceProvider.Query().
		Where(
			resourceprovider.OrganizationalUnitID(obj.ID),
		).Paginate(
		ctx,
		after,
		first,
		before,
		last,
		generated.WithResourceProviderOrder(orderBy),
		generated.WithResourceProviderFilter(where.Filter),
	)
}

// OrganizationalUnit is the resolver for the organizationalUnit field.
func (r *resourceProviderResolver) OrganizationalUnit(ctx context.Context, obj *generated.ResourceProvider) (*OrganizationalUnit, error) {
	return &OrganizationalUnit{ID: obj.OrganizationalUnitID}, nil
}

// OrganizationalUnit returns OrganizationalUnitResolver implementation.
func (r *Resolver) OrganizationalUnit() OrganizationalUnitResolver {
	return &organizationalUnitResolver{r}
}

type organizationalUnitResolver struct{ *Resolver }
