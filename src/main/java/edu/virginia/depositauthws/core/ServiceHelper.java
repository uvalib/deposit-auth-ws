package edu.virginia.depositauthws.core;

import edu.virginia.depositauthws.models.AuthDetails;
import edu.virginia.depositauthws.models.DepositDetails;
import org.apache.commons.lang3.tuple.Pair;

import javax.ws.rs.core.Response;

public class ServiceHelper {

    public static final String badIdError = "Missing/Bad Id field";
    public static final String badCidError = "Missing/Bad computing Id field";
    public static final String badLidError = "Missing/Bad document Id field";
    public static final String badDocTypeError = "Missing/Bad document type field";
    public static final String badDateError = "Missing/Bad date field";
    public static final String badAuthTokenError = "Missing/Bad authorization token";

    //
    // validate inbound request parameters
    //
    public static Pair<Response.Status, String> validateCanDepositRequest( String cid, String doctype ) {

        // check the cid
        if( !validateCid( cid ) ) {
            return( Pair.of( Response.Status.BAD_REQUEST, badCidError ) );
        }

        // check the doctype
        if( !validateDocType( doctype ) ) {
            return( Pair.of( Response.Status.BAD_REQUEST, badDocTypeError ) );
        }

        // all good
        return (Pair.of(Response.Status.OK, null));
    }

    //
    // validate inbound request parameters
    //
    public static Pair<Response.Status, String> validateDoDepositRequest( String cid, String doctype, DepositDetails details ) {

        // check the cid
        if( !validateCid( cid ) ) {
            return( Pair.of( Response.Status.BAD_REQUEST, badCidError ) );
        }

        // check the doctype
        if( !validateDocType( doctype ) ) {
            return( Pair.of( Response.Status.BAD_REQUEST, badDocTypeError ) );
        }

        // check the document Id
        if( !validateLid( details.getDocId( ) ) ) {
            return( Pair.of(Response.Status.BAD_REQUEST, badLidError ));
        }

        // check the auth token
        if( !validateAuthToken( details.getAuth( ) ) ) {
            return( Pair.of(Response.Status.UNAUTHORIZED, badAuthTokenError ));
        }

        // all good
        return (Pair.of(Response.Status.OK, null));
    }

    //
    // validate inbound request parameters
    //
    public static Pair<Response.Status, String> validateDoDeleteRequest( String id, AuthDetails details ) {

        // check the cid
        if( !validateId( id ) ) {
            return( Pair.of( Response.Status.BAD_REQUEST, badIdError ) );
        }

        // check the auth token
        if( !validateAuthToken( details.getAuth( ) ) ) {
            return( Pair.of(Response.Status.UNAUTHORIZED, badAuthTokenError ));
        }

        // all good
        return (Pair.of(Response.Status.OK, null));
    }

    //
    // validate inbound request parameters
    //
    public static Pair<Response.Status, String> validateImportRequest( String date, AuthDetails details ) {
        return (validateImportExportRequest( date, details ) );
    }

    //
    // validate inbound request parameters
    //
    public static Pair<Response.Status, String> validateExportRequest( String date, AuthDetails details ) {
        return (validateImportExportRequest( date, details ) );
    }

    //
    // validate inbound request parameters
    //
    private static Pair<Response.Status, String> validateImportExportRequest( String date, AuthDetails details ) {

        // check the auth token
        if( !validateAuthToken( details.getAuth( ) ) ) {
            return( Pair.of(Response.Status.UNAUTHORIZED, badAuthTokenError ));
        }

        // check the date
        if( !validateDate( date ) ) {
            return( Pair.of(Response.Status.BAD_REQUEST, badDateError ));
        }

        return (Pair.of(Response.Status.OK, null));
    }

    //
    // validate a record Id
    //
    private static Boolean validateId( String id ) {
        if( id.isEmpty( ) ) return( false );
        if( !id.matches( "\\d+" ) ) return( false );
        return( true );
    }

    //
    // validate a computing Id
    //
    private static Boolean validateCid( String cid ) {
        if( cid.isEmpty( ) ) return( false );
        return( true );
    }

    //
    // validate a document Id
    //
    private static Boolean validateLid( String lid ) {
        if( lid.isEmpty( ) ) return( false );
        return( true );
    }

    //
    // validate a doctype
    //
    private static Boolean validateDocType( String doctype ) {
        if( doctype.isEmpty( ) ) return( false );
        return( true );
    }

    //
    // validate an import/export date
    //
    private static Boolean validateDate( String date ) {
        if( date.isEmpty( ) ) return( false );
        if( !date.matches( "\\d{8}" ) ) return( false );
        return( true );
    }

    //
    // validate an auth token
    //
    private static Boolean validateAuthToken( String auth ) {
        if( auth.isEmpty() || !auth.equals("super-secret")) return( false );
        return( true );
    }

    public static Boolean isValid( Response.Status status ) {
        return( status.equals( Response.Status.OK ) );
    }

}
