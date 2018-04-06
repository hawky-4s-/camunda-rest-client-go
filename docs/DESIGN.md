# Design of Octoman

- State machine based on reconciler pattern for each resource and overall
- Schema for resources which maps external schema to internal (GitHub) ones
- Each resource has operations for CRUD (and exists)
- Registry for resources

## Resources

- Organizations
- Repositories
- Teams
- Users
- Membership (User / Team / Repository)

## Resource dependency graph

Organization
- Teams
- Members (Users)
- Repositories

Team
- Members

Repository
- Teams
- Collaborateur

## Resource cycle

Fetch -> Compare -> Plan -> Apply
