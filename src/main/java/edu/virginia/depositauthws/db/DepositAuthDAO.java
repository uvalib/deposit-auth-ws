package edu.virginia.depositauthws.db;

import edu.virginia.depositauthws.models.DepositAuth;
import edu.virginia.depositauthws.mapper.DepositAuthMapper;

import org.skife.jdbi.v2.sqlobject.Bind;
import org.skife.jdbi.v2.sqlobject.BindBean;
import org.skife.jdbi.v2.sqlobject.SqlQuery;
import org.skife.jdbi.v2.sqlobject.SqlUpdate;
import org.skife.jdbi.v2.sqlobject.customizers.RegisterMapper;

import java.util.List;

@RegisterMapper( DepositAuthMapper.class)
public interface DepositAuthDAO {

    // get all deposit auth records
    @SqlQuery( "select * from deposit_auth" )
    List<DepositAuth> getAll();

    // get by employee Id
    @SqlQuery( "select * from deposit_auth where eid = :eid" )
    List<DepositAuth> findByEid( @Bind("eid") String eid );

    // get by computing Id
    @SqlQuery( "select * from deposit_auth where cid = :cid" )
    List<DepositAuth> findByCid( @Bind("cid") String cid );

    // get by libra Id
    @SqlQuery( "select * from deposit_auth where lid = :lid" )
    List<DepositAuth> findByLid( @Bind("lid") String lid );

    // get by computing Id and doctype
    @SqlQuery( "select * from deposit_auth where cid = :cid AND doctype = :doctype" )
    List<DepositAuth> findByCidAndDoctype( @Bind("cid") String cid, @Bind("doctype") String doctype );

    // get auth records suitable for export
    @SqlQuery( "select * from deposit_auth where accepted_at IS NOT NULL AND exported_at IS NULL" )
    List<DepositAuth> getForExport( );

    // add a new item
    @SqlUpdate( "insert into deposit_auth (id, eid, cid, first_name, middle_name, last_name, career, program, plan, degree, title, doctype, lid, approved_at, created_at) values (0, :eid, :cid, :firstName, :middleName, :lastName, :career, :program, :plan, :degree, :title, :doctype, :lid, :approvedAt, NOW( ))" )
    int insert( @BindBean DepositAuth depositAuth );

    // update an existing item
    @SqlUpdate( "update deposit_auth set cid = :cid, first_name = :firstName, middle_name = :middleName, last_name = :lastName, title = :title, degree = :degree, approved_at = :approvedAt, updated_at = NOW( ) where id = :id" )
    int update( @BindBean DepositAuth depositAuth );

    // update the exported timestamp
    @SqlUpdate( "update deposit_auth set exported_at = NOW( ), updated_at = NOW( ) where id = :id" )
    int markExported(  @Bind("id") String id );

    // delete an auth record
    @SqlUpdate( "delete from deposit_auth where id = :id" )
    int delete(  @Bind("id") String id );
}