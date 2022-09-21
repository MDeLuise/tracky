package com.github.mdeluise.tracky.common;

import org.modelmapper.ModelMapper;

public abstract class AbstractDTOConverter<T, D> {
    protected final ModelMapper modelMapper;


    protected AbstractDTOConverter(ModelMapper modelMapper) {
        this.modelMapper = modelMapper;
    }


    public abstract T convertFromDTO(D dto);

    public abstract D convertToDTO(T data);
}
