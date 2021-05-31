// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"italking.tomotoes.com/m/v1/ent/notification"
	"italking.tomotoes.com/m/v1/ent/predicate"
	"italking.tomotoes.com/m/v1/ent/user"
)

// NotificationUpdate is the builder for updating Notification entities.
type NotificationUpdate struct {
	config
	hooks    []Hook
	mutation *NotificationMutation
}

// Where adds a new predicate for the NotificationUpdate builder.
func (nu *NotificationUpdate) Where(ps ...predicate.Notification) *NotificationUpdate {
	nu.mutation.predicates = append(nu.mutation.predicates, ps...)
	return nu
}

// SetContent sets the "content" field.
func (nu *NotificationUpdate) SetContent(s string) *NotificationUpdate {
	nu.mutation.SetContent(s)
	return nu
}

// SetType sets the "type" field.
func (nu *NotificationUpdate) SetType(n notification.Type) *NotificationUpdate {
	nu.mutation.SetType(n)
	return nu
}

// SetCreatedAt sets the "created_at" field.
func (nu *NotificationUpdate) SetCreatedAt(t time.Time) *NotificationUpdate {
	nu.mutation.SetCreatedAt(t)
	return nu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableCreatedAt(t *time.Time) *NotificationUpdate {
	if t != nil {
		nu.SetCreatedAt(*t)
	}
	return nu
}

// SetUpdatedAt sets the "updated_at" field.
func (nu *NotificationUpdate) SetUpdatedAt(t time.Time) *NotificationUpdate {
	nu.mutation.SetUpdatedAt(t)
	return nu
}

// SetSenderID sets the "sender" edge to the User entity by ID.
func (nu *NotificationUpdate) SetSenderID(id string) *NotificationUpdate {
	nu.mutation.SetSenderID(id)
	return nu
}

// SetSender sets the "sender" edge to the User entity.
func (nu *NotificationUpdate) SetSender(u *User) *NotificationUpdate {
	return nu.SetSenderID(u.ID)
}

// SetReceiverID sets the "receiver" edge to the User entity by ID.
func (nu *NotificationUpdate) SetReceiverID(id string) *NotificationUpdate {
	nu.mutation.SetReceiverID(id)
	return nu
}

// SetReceiver sets the "receiver" edge to the User entity.
func (nu *NotificationUpdate) SetReceiver(u *User) *NotificationUpdate {
	return nu.SetReceiverID(u.ID)
}

// Mutation returns the NotificationMutation object of the builder.
func (nu *NotificationUpdate) Mutation() *NotificationMutation {
	return nu.mutation
}

// ClearSender clears the "sender" edge to the User entity.
func (nu *NotificationUpdate) ClearSender() *NotificationUpdate {
	nu.mutation.ClearSender()
	return nu
}

// ClearReceiver clears the "receiver" edge to the User entity.
func (nu *NotificationUpdate) ClearReceiver() *NotificationUpdate {
	nu.mutation.ClearReceiver()
	return nu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NotificationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	nu.defaults()
	if len(nu.hooks) == 0 {
		if err = nu.check(); err != nil {
			return 0, err
		}
		affected, err = nu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NotificationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nu.check(); err != nil {
				return 0, err
			}
			nu.mutation = mutation
			affected, err = nu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nu.hooks) - 1; i >= 0; i-- {
			mut = nu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NotificationUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NotificationUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NotificationUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nu *NotificationUpdate) defaults() {
	if _, ok := nu.mutation.UpdatedAt(); !ok {
		v := notification.UpdateDefaultUpdatedAt()
		nu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nu *NotificationUpdate) check() error {
	if v, ok := nu.mutation.Content(); ok {
		if err := notification.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf("ent: validator failed for field \"content\": %w", err)}
		}
	}
	if v, ok := nu.mutation.GetType(); ok {
		if err := notification.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if _, ok := nu.mutation.SenderID(); nu.mutation.SenderCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"sender\"")
	}
	if _, ok := nu.mutation.ReceiverID(); nu.mutation.ReceiverCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"receiver\"")
	}
	return nil
}

func (nu *NotificationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   notification.Table,
			Columns: notification.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: notification.FieldID,
			},
		},
	}
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notification.FieldContent,
		})
	}
	if value, ok := nu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: notification.FieldType,
		})
	}
	if value, ok := nu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: notification.FieldCreatedAt,
		})
	}
	if value, ok := nu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: notification.FieldUpdatedAt,
		})
	}
	if nu.mutation.SenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notification.SenderTable,
			Columns: []string{notification.SenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.SenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notification.SenderTable,
			Columns: []string{notification.SenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nu.mutation.ReceiverCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notification.ReceiverTable,
			Columns: []string{notification.ReceiverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.ReceiverIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notification.ReceiverTable,
			Columns: []string{notification.ReceiverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notification.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// NotificationUpdateOne is the builder for updating a single Notification entity.
type NotificationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NotificationMutation
}

// SetContent sets the "content" field.
func (nuo *NotificationUpdateOne) SetContent(s string) *NotificationUpdateOne {
	nuo.mutation.SetContent(s)
	return nuo
}

// SetType sets the "type" field.
func (nuo *NotificationUpdateOne) SetType(n notification.Type) *NotificationUpdateOne {
	nuo.mutation.SetType(n)
	return nuo
}

// SetCreatedAt sets the "created_at" field.
func (nuo *NotificationUpdateOne) SetCreatedAt(t time.Time) *NotificationUpdateOne {
	nuo.mutation.SetCreatedAt(t)
	return nuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableCreatedAt(t *time.Time) *NotificationUpdateOne {
	if t != nil {
		nuo.SetCreatedAt(*t)
	}
	return nuo
}

// SetUpdatedAt sets the "updated_at" field.
func (nuo *NotificationUpdateOne) SetUpdatedAt(t time.Time) *NotificationUpdateOne {
	nuo.mutation.SetUpdatedAt(t)
	return nuo
}

// SetSenderID sets the "sender" edge to the User entity by ID.
func (nuo *NotificationUpdateOne) SetSenderID(id string) *NotificationUpdateOne {
	nuo.mutation.SetSenderID(id)
	return nuo
}

// SetSender sets the "sender" edge to the User entity.
func (nuo *NotificationUpdateOne) SetSender(u *User) *NotificationUpdateOne {
	return nuo.SetSenderID(u.ID)
}

// SetReceiverID sets the "receiver" edge to the User entity by ID.
func (nuo *NotificationUpdateOne) SetReceiverID(id string) *NotificationUpdateOne {
	nuo.mutation.SetReceiverID(id)
	return nuo
}

// SetReceiver sets the "receiver" edge to the User entity.
func (nuo *NotificationUpdateOne) SetReceiver(u *User) *NotificationUpdateOne {
	return nuo.SetReceiverID(u.ID)
}

// Mutation returns the NotificationMutation object of the builder.
func (nuo *NotificationUpdateOne) Mutation() *NotificationMutation {
	return nuo.mutation
}

// ClearSender clears the "sender" edge to the User entity.
func (nuo *NotificationUpdateOne) ClearSender() *NotificationUpdateOne {
	nuo.mutation.ClearSender()
	return nuo
}

// ClearReceiver clears the "receiver" edge to the User entity.
func (nuo *NotificationUpdateOne) ClearReceiver() *NotificationUpdateOne {
	nuo.mutation.ClearReceiver()
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NotificationUpdateOne) Select(field string, fields ...string) *NotificationUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Notification entity.
func (nuo *NotificationUpdateOne) Save(ctx context.Context) (*Notification, error) {
	var (
		err  error
		node *Notification
	)
	nuo.defaults()
	if len(nuo.hooks) == 0 {
		if err = nuo.check(); err != nil {
			return nil, err
		}
		node, err = nuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NotificationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nuo.check(); err != nil {
				return nil, err
			}
			nuo.mutation = mutation
			node, err = nuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nuo.hooks) - 1; i >= 0; i-- {
			mut = nuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NotificationUpdateOne) SaveX(ctx context.Context) *Notification {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NotificationUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NotificationUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nuo *NotificationUpdateOne) defaults() {
	if _, ok := nuo.mutation.UpdatedAt(); !ok {
		v := notification.UpdateDefaultUpdatedAt()
		nuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nuo *NotificationUpdateOne) check() error {
	if v, ok := nuo.mutation.Content(); ok {
		if err := notification.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf("ent: validator failed for field \"content\": %w", err)}
		}
	}
	if v, ok := nuo.mutation.GetType(); ok {
		if err := notification.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if _, ok := nuo.mutation.SenderID(); nuo.mutation.SenderCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"sender\"")
	}
	if _, ok := nuo.mutation.ReceiverID(); nuo.mutation.ReceiverCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"receiver\"")
	}
	return nil
}

func (nuo *NotificationUpdateOne) sqlSave(ctx context.Context) (_node *Notification, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   notification.Table,
			Columns: notification.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: notification.FieldID,
			},
		},
	}
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Notification.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notification.FieldID)
		for _, f := range fields {
			if !notification.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != notification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notification.FieldContent,
		})
	}
	if value, ok := nuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: notification.FieldType,
		})
	}
	if value, ok := nuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: notification.FieldCreatedAt,
		})
	}
	if value, ok := nuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: notification.FieldUpdatedAt,
		})
	}
	if nuo.mutation.SenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notification.SenderTable,
			Columns: []string{notification.SenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.SenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notification.SenderTable,
			Columns: []string{notification.SenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nuo.mutation.ReceiverCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notification.ReceiverTable,
			Columns: []string{notification.ReceiverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.ReceiverIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notification.ReceiverTable,
			Columns: []string{notification.ReceiverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Notification{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notification.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}