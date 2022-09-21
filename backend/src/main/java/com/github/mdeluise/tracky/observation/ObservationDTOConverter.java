package com.github.mdeluise.tracky.observation;

import com.github.mdeluise.tracky.common.AbstractDTOConverter;
import org.modelmapper.ModelMapper;
import org.springframework.stereotype.Component;

import java.util.Date;

@Component
public class ObservationDTOConverter extends AbstractDTOConverter<Observation, ObservationDTO> {


    public ObservationDTOConverter(ModelMapper modelMapper) {
        super(modelMapper);
    }


    @Override
    public Observation convertFromDTO(ObservationDTO dto) {
        final Observation result = modelMapper.map(dto, Observation.class);
        if (result.getInstant() == null) {
            result.setInstant(new Date());
        }
        return result;
    }


    @Override
    public ObservationDTO convertToDTO(Observation data) {
        return modelMapper.map(data, ObservationDTO.class);
    }
}
