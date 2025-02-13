package com.example.profile_service.models;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.Table;
import lombok.Data;

@Entity
@Table(name = "profile")
@Data
public class Profile {
  @Id
  @Column(name = "user_id", nullable = false)
  private String userId;

  @Column(unique = true, nullable = false)
  private String nickname;

  @Column
  private String country;

  @Column(columnDefinition = "TEXT")
  private String description;
}
