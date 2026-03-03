# Contributing

## Release Management

### Nightly Git Tag

There is a nightly job that runs and tags the latest commit IF there has been a
commit that day.

## Troubleshooting

If you see a local git error

```bash
git fetch --prune --tags
From https://codeberg.org/yaadata/bina
 ! [rejected]        nightly    -> nightly  (would clobber existing tag)
```

This can be resolved by running

```bash
just ref-tags
```
