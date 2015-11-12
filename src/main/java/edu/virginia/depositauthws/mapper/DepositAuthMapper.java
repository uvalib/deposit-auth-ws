package edu.virginia.depositauthws.mapper;

import edu.virginia.depositauthws.models.DepositAuth;

import org.skife.jdbi.v2.StatementContext;
import org.skife.jdbi.v2.tweak.ResultSetMapper;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Timestamp;
import java.text.SimpleDateFormat;

public class DepositAuthMapper implements ResultSetMapper<DepositAuth> {

    SimpleDateFormat sdf = new SimpleDateFormat( "yyyy-MM-dd HH:mm:ss z" );

    public DepositAuth map( int index, ResultSet resultSet, StatementContext statementContext ) throws SQLException {
        return new DepositAuth( )
           .setId( resultSet.getLong( "id" ) )
           .setCid( resultSet.getString( "cid" ) )
           .setDoctype( resultSet.getString( "doctype" ) )
           .setLid( resultSet.getString( "lid" ) )
           .setApprovedAt( formatDate( resultSet.getTimestamp( "libra_approved_at" ) ) )
           .setExportedAt( formatDate( resultSet.getTimestamp( "exported_at" ) ) )
           .setCreatedAt( formatDate( resultSet.getTimestamp( "created_at" ) ) )
           .setUpdatedAt( formatDate( resultSet.getTimestamp( "updated_at" ) ) );
    }

    private String formatDate( Timestamp datetime ) {
       if( datetime != null ) {
           return( sdf.format( datetime ) );
       }
       return( "" );
    }
}