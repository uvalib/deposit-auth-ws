package edu.virginia.depositauthws.resources;

import edu.virginia.depositauthws.models.Saying;
import edu.virginia.depositauthws.models.SingleStatus;
import edu.virginia.depositauthws.models.AllStatus;

import com.google.common.base.Optional;
import com.codahale.metrics.annotation.Timed;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.QueryParam;
import javax.ws.rs.core.MediaType;
import java.util.concurrent.atomic.AtomicLong;

@Path( "/sis" )
@Produces( MediaType.APPLICATION_JSON )
public class ServiceResource {

    private final String template;
    private final String defaultName;
    private final AtomicLong counter;

    public ServiceResource( String template, String defaultName ) {
        this.template = template;
        this.defaultName = defaultName;
        this.counter = new AtomicLong( );
    }

    @GET
    @Path( "/hello" )
    @Timed
    public Saying sayHello( @QueryParam( "name" ) Optional<String> name ) {
        final String value = String.format( template, name.or( defaultName ) );
        return new Saying( counter.incrementAndGet( ), value );
    }

    @GET
    @Path( "/status" )
    @Timed
    public SingleStatus singleStatus( @QueryParam( "name" ) Optional<String> name ) {
        final String value = String.format( template, name.or( defaultName ) );
        return new SingleStatus( counter.incrementAndGet( ), value );
    }

    @GET
    @Path( "/all" )
    @Timed
    public AllStatus allStatus( @QueryParam( "name" ) Optional<String> name ) {
        final String value = String.format( template, name.or( defaultName ) );
        return new AllStatus( counter.incrementAndGet( ), value );
    }

}