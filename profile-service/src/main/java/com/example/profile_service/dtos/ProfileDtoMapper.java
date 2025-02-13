package com.example.profile_service.dtos;

import com.example.profile_service.models.Profile;
import org.mapstruct.Mapper;
import org.mapstruct.MappingTarget;
import org.mapstruct.factory.Mappers;

@Mapper
public interface ProfileDtoMapper {
  ProfileDtoMapper INSTANCE = Mappers.getMapper(ProfileDtoMapper.class);

  Profile createProfileDtoToProfile(CreateProfileDto dto);

  SearchProfileDto profileToSearchProfileDto(Profile p);

  Profile updateProfile(@MappingTarget Profile p, UpdateProfileDto dto);

}
