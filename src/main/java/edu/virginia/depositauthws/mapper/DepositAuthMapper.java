package edu.virginia.depositauthws.mapper;

import edu.virginia.depositauthws.models.DepositAuth;

import org.skife.jdbi.v2.StatementContext;
import org.skife.jdbi.v2.tweak.ResultSetMapper;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Timestamp;
import java.sql.Date;
import java.text.SimpleDateFormat;

public class DepositAuthMapper implements ResultSetMapper<DepositAuth> {

    public static final SimpleDateFormat datetimeFormat = new SimpleDateFormat( "yyyy-MM-dd HH:mm:ss z" );
    public static final SimpleDateFormat dateFormat = new SimpleDateFormat( "yyyy-MM-dd" );

    public DepositAuth map( int index, ResultSet resultSet, StatementContext statementContext ) throws SQLException {
        return new DepositAuth( )
           .setId( resultSet.getLong( "id" ) )
           .setEid( resultSet.getString( "eid" ) )
           .setCid( resultSet.getString( "cid" ) )

           .setFirstName( resultSet.getString( "first_name" ) )
           .setMiddleName( resultSet.getString( "middle_name" ) )
           .setLastName( resultSet.getString( "last_name" ) )

           .setCareer( resultSet.getString( "career" ) )
           .setProgram( resultSet.getString( "program" ) )
           .setPlan( resultSet.getString( "plan" ) )
           .setDegree( resultSet.getString( "degree" ) )

           .setTitle( resultSet.getString( "title" ) )
           .setDoctype( resultSet.getString( "doctype" ) )
           .setLid( resultSet.getString( "lid" ) )

           .setApprovedAt( formatDate( resultSet.getDate( "approved_at" ) ) )
           .setAcceptedAt( formatDateTime( resultSet.getTimestamp( "accepted_at" ) ) )
           .setExportedAt( formatDateTime( resultSet.getTimestamp( "exported_at" ) ) )
           .setCreatedAt( formatDateTime( resultSet.getTimestamp( "created_at" ) ) )
           .setUpdatedAt( formatDateTime( resultSet.getTimestamp( "updated_at" ) ) );
    }

    private String formatDate( Date date ) {
       if( date != null ) {
           return( dateFormat.format( date ) );
       }
       return( "" );
    }

    private String formatDateTime( Timestamp datetime ) {
        if( datetime != null ) {
            return( datetimeFormat.format( datetime ) );
        }
        return( "" );
    }
}