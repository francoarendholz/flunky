# Description: Updates the version string in the VERSION file
# Usage: updateVersion.sh [major|minor|patch]
# Example: updateVersion.sh minor

# Check that a part was specified
if [ -z "$1" ]; then
    # If no part was specified, default to patch
    part="patch"
else
    # If a part was specified, set it to the part variable
    part="$1"
fi

# Check that the specified part is valid
if [ "$part" != "major" ] && [ "$part" != "minor" ] && [ "$part" != "patch" ]; then
    echo "Invalid part specified"
    exit 1
fi

# Read the version string from the VERSION file
version=$(cat VERSION)

# Split the version string into an array
IFS='.' read -ra versionArray <<< "$version"

# Increment the specified part of the version string
if [ "$part" == "major" ]; then
    versionArray[0]=$((${versionArray[0]} + 1))
elif [ "$part" == "minor" ]; then
    versionArray[1]=$((${versionArray[1]} + 1))
elif [ "$part" == "patch" ]; then
    versionArray[2]=$((${versionArray[2]} + 1))
fi

# Join the version array into a string
version="${versionArray[0]}.${versionArray[1]}.${versionArray[2]}"
echo $version

# Write the version string to the VERSION file
echo $version > VERSION
