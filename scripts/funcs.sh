# funcs and variables to make curling strava api easier
# instructions:
#  source scripts/funcs.sh <stravaToken>
# be sure to provide strava token

stravaApi="https://www.strava.com/api/v3"
stravaAuthStr="Authorization: Bearer $1"

#retrieves a segment by id
retrieveSegment() {
   echo "$#"
   if [ "$#" -ne 1 ]; then
      echo "usage: retrieveSegment id"
      exit 1
   fi

   local segId="$1"

   curl -G "$stravaApi/segments/$segId" -H "$stravaAuthStr"
}

# fetches segment efforts for given segment id
listEfforts() {
   local segId="$1"

   curl -G "$stravaApi/segments/$segId/all_efforts" -H "$stravaAuthStr"
}

# fetches athlete with given id
retrieveAthlete() {
   local athleteId="$1"

   curl -G "$stravaApi/athletes/$athleteId" -H "$stravaAuthStr"
}
