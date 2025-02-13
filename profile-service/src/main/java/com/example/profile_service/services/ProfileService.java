package com.example.profile_service.services;

import com.example.profile_service.dtos.CreateProfileDto;
import com.example.profile_service.dtos.ProfileDtoMapper;
import com.example.profile_service.dtos.SearchProfileDto;
import com.example.profile_service.dtos.UpdateProfileDto;
import com.example.profile_service.exceptions.ProfileAlreadyExistsException;
import com.example.profile_service.exceptions.ProfileNotFoundException;
import com.example.profile_service.models.Profile;
import com.example.profile_service.repositories.ProfileRepository;
import org.springframework.stereotype.Service;
import java.util.List;

@Service
public class ProfileService {
  private ProfileRepository profileRepository;

  public ProfileService(ProfileRepository profileRepository) {
    this.profileRepository = profileRepository;
  }

  public Profile getProfile(String id) throws ProfileNotFoundException {
    return this.profileRepository.findById(id).orElseThrow(() -> new ProfileNotFoundException(id));
  }

  public List<SearchProfileDto> searchProfileByNickname(String query) {
    return this.profileRepository.searchByNickname(query).stream().map(ProfileDtoMapper.INSTANCE::profileToSearchProfileDto).toList();
  }

  public Profile createProfile(CreateProfileDto profileDto, String userId) throws Exception {
    if (this.profileRepository.findById(userId).isPresent()) {
      throw new ProfileAlreadyExistsException();
    }
    Profile profile = ProfileDtoMapper.INSTANCE.createProfileDtoToProfile(profileDto);
    profile.setUserId(userId);

    return this.profileRepository.save(profile);
  }

  public Profile updateProfile(String id, UpdateProfileDto profile) throws ProfileNotFoundException {
    Profile p = this.profileRepository.findById(id).orElseThrow(() -> new ProfileNotFoundException(id));
    Profile updated = ProfileDtoMapper.INSTANCE.updateProfile(p, profile);
    return this.profileRepository.save(updated);
  }

  public void deleteProfile(String id) throws ProfileNotFoundException {
    this.profileRepository.findById(id).orElseThrow(() -> new ProfileNotFoundException(id));
    this.profileRepository.deleteById(id);
  }
}
