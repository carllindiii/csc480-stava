import javastrava.api.v3.auth.AuthorisationService;
import javastrava.api.v3.auth.impl.retrofit.AuthorisationServiceImpl;
import javastrava.api.v3.auth.model.Token;
import javastrava.api.v3.model.StravaSegment;
import javastrava.api.v3.service.Strava;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/**
 * Created by carllindiii on 10/15/15.
 */
public class Main {
    public static void main(String[] args) throws Exception{
        // Authenticate with Strava
        AuthorisationService service = new AuthorisationServiceImpl();
        AuthenticationInfo authInfo = new AuthenticationInfo();

        // get code using oauth


//        Token token = service.tokenExchange(authInfo.getApplicationClientID(), authInfo.getClientSecret(), "");
        Token token = new Token();
        Strava strava = new Strava(token);

        // list of segments
        //10 Segment IDs: 365235, 6452581, 664647, 1089563, 4956199, 2187, 5732938, 654030, 616554, 3139189
        List<Integer> segmentIDs = Arrays.asList(365235, 6452581, 664647, 1089563, 4956199, 2187, 5732938, 654030, 616554, 3139189);
        List<StravaSegment> segments = new ArrayList<>();

        // get all strava segments (not asynchronous)
        for(Integer segmentID : segmentIDs) {
            segments.add(strava.getSegment(segmentID));
        }

        for(StravaSegment segment : segments) {
            System.out.println("Segment ID: " + segment.getId() + " segment type: " + segment.getActivityType() + "\n");
        }

    }
}
