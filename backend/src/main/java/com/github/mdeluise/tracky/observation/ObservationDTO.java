package com.github.mdeluise.tracky.observation;

import com.fasterxml.jackson.annotation.JsonProperty;
import io.swagger.v3.oas.annotations.media.Schema;

import java.util.Date;
import java.util.Objects;

@Schema(name = "Observation", description = "Represents an observation.")
public class ObservationDTO {
    @Schema(description = "ID of the observation.", accessMode = Schema.AccessMode.READ_ONLY)
    private Long id;
    @Schema(description = "ID of the tracker that hold this observation.")
    private Long trackerId;
    @Schema(description = "Name of the tracker that hold this observation.", accessMode = Schema.AccessMode.READ_ONLY)
    private String trackerName;
    @Schema(description = "Note of the observation.")
    private String note;
    @Schema(
        name = "unit", description = "Unit of measure of the tracker that hold this observation.",
        accessMode = Schema.AccessMode.READ_ONLY
    )
    @JsonProperty("unit")
    private String trackerUnit;
    @Schema(description = "Date of the observation.")
    private Date instant;
    @Schema(description = "Value of the observation.", example = "42.24")
    private float value;


    public Long getId() {
        return id;
    }


    public void setId(Long id) {
        this.id = id;
    }


    public Long getTrackerId() {
        return trackerId;
    }


    public void setTrackerId(Long trackerId) {
        this.trackerId = trackerId;
    }


    public String getTrackerName() {
        return trackerName;
    }


    public void setTrackerName(String trackerName) {
        this.trackerName = trackerName;
    }


    public void setTrackerUnit(String trackerUnit) {
        this.trackerUnit = trackerUnit;
    }


    public Date getInstant() {
        return instant;
    }


    public void setInstant(Date instant) {
        this.instant = instant;
    }


    public Number getValue() {
        return value;
    }


    public void setValue(Number value) {
        this.value = value.floatValue();
    }


    public String getNote() {
        return note;
    }


    public void setNote(String note) {
        this.note = note;
    }


    @Override
    public boolean equals(Object o) {
        if (this == o) {
            return true;
        }
        if (o == null || getClass() != o.getClass()) {
            return false;
        }
        ObservationDTO that = (ObservationDTO) o;
        return Objects.equals(id, that.id);
    }


    @Override
    public int hashCode() {
        return Objects.hash(id);
    }
}
