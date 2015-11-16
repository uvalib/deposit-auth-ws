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

    @SqlQuery( "select * from deposit_auth" )
    List<DepositAuth> getAll();

    @SqlQuery( "select * from deposit_auth where cid = :cid" )
    List<DepositAuth> findByCid( @Bind("cid") String cid );

    @SqlQuery( "select * from deposit_auth where lid = :lid" )
    List<DepositAuth> findByLid( @Bind("lid") String lid );

    @SqlQuery( "select * from deposit_auth where cid = :cid AND doctype = :doctype" )
    List<DepositAuth> findByCidAndDoctype( @Bind("cid") String cid, @Bind("doctype") String doctype );

    @SqlQuery( "select * from deposit_auth where approved_at IS NOT NULL AND exported_at IS NULL" )
    List<DepositAuth> getForExport( );

    //@SqlUpdate( "update deposit_auth set lid = :lid where CID = :cid" )
    //int update(@BindBean Person person);

    //@SqlUpdate( "insert into deposit_auth (id, cid, lid) values (:id, :cid, :lid)" )
    //int insert( @BindBean DepositAuth depositAuth );
}