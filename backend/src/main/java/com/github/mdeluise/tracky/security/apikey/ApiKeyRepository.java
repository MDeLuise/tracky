package com.github.mdeluise.tracky.security.apikey;

import com.github.mdeluise.tracky.authentication.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.Collection;
import java.util.Optional;

@Repository
public interface ApiKeyRepository extends JpaRepository<ApiKey, Long> {

    Collection<ApiKey> findByUser(User user);

    Optional<ApiKey> findByValue(String apiKeyValue);

    boolean existsByValue(String apiKeyValue);

    Optional<ApiKey> findByUserAndName(User user, String name);
}
