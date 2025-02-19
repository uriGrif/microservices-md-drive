package com.example.profile_service.models;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.PrePersist;
import jakarta.persistence.PreUpdate;
import jakarta.persistence.Table;
import jakarta.persistence.Temporal;
import jakarta.persistence.TemporalType;
import lombok.Data;
import java.util.Date;

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

  @Column(name = "created_at")
  @Temporal(TemporalType.TIMESTAMP)
  private Date createdAt;

  @Column(name = "updated_at")
  @Temporal(TemporalType.TIMESTAMP)
  private Date updatedAt;

  @PrePersist
  public void setCreatedTime() {
    this.createdAt = new Date();
  }

  @PreUpdate
  public void setUpdatedTime() {
    this.updatedAt = new Date();
  }
}
