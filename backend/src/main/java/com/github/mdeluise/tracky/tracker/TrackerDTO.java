package com.github.mdeluise.tracky.tracker;

import com.github.mdeluise.tracky.observation.ObservationDTO;
import io.swagger.v3.oas.annotations.media.Schema;

import java.util.Objects;

@Schema(name = "Tracker", description = "Represents a tracker.")
public class TrackerDTO {
    @Schema(description = "ID of the tracker.", accessMode = Schema.AccessMode.READ_ONLY)
    private Long id;
    @Schema(description = "Name of the tracker.", example = "GPL price")
    private String name;
    @Schema(description = "Description of the tracker.", example = "Track the GPL price over time")
    private String description;
    @Schema(description = "Unit of measurement of the tracker.", example = "€/L")
    private String unit;
    @Schema(
        description = "Last observation inserted in the tracker.", accessMode = Schema.AccessMode.READ_ONLY
    )
    private ObservationDTO lastObservation;


    public Long getId() {
        return id;
    }


    public void setId(Long id) {
        this.id = id;
    }


    public String getName() {
        return name;
    }


    public void setName(String name) {
        this.name = name;
    }


    public String getDescription() {
        return description;
    }


    public void setDescription(String description) {
        this.description = description;
    }


    public String getUnit() {
        return unit;
    }


    public void setUnit(String unit) {
        this.unit = unit;
    }


    public void setLastObservation(ObservationDTO lastObservation) {
        this.lastObservation = lastObservation;
    }


    public ObservationDTO getLastObservation() {
        return lastObservation;
    }


    @Override
    public boolean equals(Object o) {
        if (this == o) {
            return true;
        }
        if (o == null || getClass() != o.getClass()) {
            return false;
        }
        TrackerDTO that = (TrackerDTO) o;
        return Objects.equals(id, that.id);
    }


    @Override
    public int hashCode() {
        return Objects.hash(id);
    }
}
