package com.example.profile_service.repositories;

import com.example.profile_service.models.Profile;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import java.util.List;

public interface ProfileRepository extends JpaRepository<Profile, String> {
  @Query(value = "SELECT * FROM profile WHERE nickname LIKE :query%", nativeQuery = true)
  List<Profile> searchByNickname(@Param("query") String query);
}
