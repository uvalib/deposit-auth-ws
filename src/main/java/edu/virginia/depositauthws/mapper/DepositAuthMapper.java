package edu.virginia.depositauthws.mapper;

import edu.virginia.depositauthws.models.DepositAuth;

import org.skife.jdbi.v2.StatementContext;
import org.skife.jdbi.v2.tweak.ResultSetMapper;

import java.sql.ResultSet;
import java.sql.SQLException;

public class DepositAuthMapper implements ResultSetMapper<DepositAuth> {
    public DepositAuth map( int index, ResultSet resultSet, StatementContext statementContext ) throws SQLException {
        return new DepositAuth( )
           .setId( resultSet.getLong( "id" ) )
           .setCid( resultSet.getString( "cid" ) )
           .setDoctype( resultSet.getString( "doctype" ) )
           .setLid( resultSet.getString( "lid" ) );
    }
}