package edu.virginia.depositauthws.models;

import com.fasterxml.jackson.annotation.JsonProperty;
import javax.validation.constraints.NotNull;

public class DepositAuth {

    @NotNull
    @JsonProperty
    private Long id;

    @NotNull
    @JsonProperty
    private String cid;

    @NotNull
    @JsonProperty
    private String doctype;

    @JsonProperty
    private String lid;

    @JsonProperty
    private String title;

    @JsonProperty
    private String program;

    @JsonProperty
    private String approvedAt;

    @JsonProperty
    private String exportedAt;

    //@NotNull
    @JsonProperty
    private String createdAt;

    @JsonProperty
    private String updatedAt;

    public DepositAuth() {
        // Jackson deserialization
    }

    //
    // getters...
    //

    public Long getId() {
        return id;
    }

    public String getCid() {
        return cid;
    }

    public String getDoctype() {
        return doctype;
    }

    public String getLid() {
        return valueOrEmptyString( lid );
    }

    public String getTitle() { return valueOrEmptyString( title ); }

    public String getProgram() {
        return valueOrEmptyString( program );
    }

    public String getApprovedAt() {
        return approvedAt;
    }

    public String getExportedAt() {
        return exportedAt;
    }

    public String getCreatedAt() {
        return createdAt;
    }

    public String getUpdatedAt() {
        return updatedAt;
    }

    //
    // setters...
    //

    public DepositAuth setId( Long id ) {
        this.id = id;
        return this;
    }

    public DepositAuth setCid( String cid ) {
        this.cid = cid;
        return this;
    }

    public DepositAuth setDoctype( String doctype ) {
        this.doctype = doctype;
        return this;
    }

    public DepositAuth setLid( String lid ) {
        this.lid = lid;
        return this;
    }

    public DepositAuth setTitle( String title ) {
        this.title = title;
        return this;
    }

    public DepositAuth setProgram( String program ) {
        this.program = program;
        return this;
    }

    public DepositAuth setApprovedAt( String approvedAt ) {
        this.approvedAt = approvedAt;
        return this;
    }

    public DepositAuth setExportedAt( String exportedAt ) {
        this.exportedAt = exportedAt;
        return this;
    }

    public DepositAuth setCreatedAt( String createdAt ) {
        this.createdAt = createdAt;
        return this;
    }

    public DepositAuth setUpdatedAt( String updatedAt ) {
        this.updatedAt = updatedAt;
        return this;
    }

    private String valueOrEmptyString( String value ) {
        return value == null ? "" : value;
    }
}
