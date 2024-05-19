#! /usr/bin/bash

VERSION=""

# get parameters
while getopts v: flag
do
  case "${flag}" in
    v) VERSION=${OPTARG};;
  esac
done

# fetch all tags in case of a shallow clone
git fetch --prune --unshallow 2>/dev/null || git fetch --prune --tags
CURRENT_VERSION=$(git describe --abbrev=0 --tags 2>/dev/null)

# set default version if no tags are present
if [[ -z $CURRENT_VERSION ]]; then
  CURRENT_VERSION='v0.1.0'
fi
echo "Current Version: $CURRENT_VERSION"

# replace . with space so can split into an array
CURRENT_VERSION_PARTS=(${CURRENT_VERSION//./ })

# get number parts, remove the initial 'v' from the major version part
VNUM1=${CURRENT_VERSION_PARTS[0]#v}
VNUM2=${CURRENT_VERSION_PARTS[1]}
VNUM3=${CURRENT_VERSION_PARTS[2]}

# increment version numbers based on the specified version type
if [[ $VERSION == 'major' ]]; then
  VNUM1=$((VNUM1 + 1))
  VNUM2=0
  VNUM3=0
elif [[ $VERSION == 'minor' ]]; then
  VNUM2=$((VNUM2 + 1))
  VNUM3=0
elif [[ $VERSION == 'patch' ]]; then
  VNUM3=$((VNUM3 + 1))
else
  echo "No version type (https://semver.org/) or incorrect type specified, try: -v [major, minor, patch]"
  exit 1
fi

# create new tag
NEW_TAG="v$VNUM1.$VNUM2.$VNUM3"
echo "($VERSION) updating $CURRENT_VERSION to $NEW_TAG"

# set up git to use the GITHUB_TOKEN for authentication
git config --global user.email "github-actions[bot]@users.noreply.github.com"
git config --global user.name "github-actions[bot]"
git remote set-url origin https://x-access-token:${GITHUB_TOKEN}@github.com/${GITHUB_REPOSITORY}

# get current hash and see if it already has a tag
GIT_COMMIT=$(git rev-parse HEAD)
NEEDS_TAG=$(git describe --contains $GIT_COMMIT 2>/dev/null)

# only tag if no tag already
if [[ -z $NEEDS_TAG ]]; then
  echo "Tagged with $NEW_TAG"
  git tag $NEW_TAG
  git push --tags
  git push
else
  echo "Already a tag on this commit"
fi

# output the new tag
echo "::set-output name=git_tag::$NEW_TAG"

exit 0