package edu.virginia.depositauthws.sis;

import org.apache.commons.lang3.tuple.Pair;

import javax.ws.rs.core.Response;

public class SisHelper {


    //
    // write any pending records to SIS
    //
    public static Pair<Response.Status, Integer> exportToSis(String fs, String date ) {
        return( Pair.of( Response.Status.OK, 0 ) );
    }

    //
    // write any pending records to SIS
    //
    public static Pair<Response.Status, Integer> importFromSis( String fs, String date ) {
        return( Pair.of( Response.Status.OK, 0 ) );
    }

    //
    // generate the input from SIS filename
    //
    private static String inputFile( String date ) {
       return( "UV_Libra_From_SIS" + date + ".txt" );
    }

    //
    // generate the output to SIS filename
    //
    private static String outputFile( String date ) {
        return( "UV_LIBRA_IN" + date + ".txt" );
    }
}
