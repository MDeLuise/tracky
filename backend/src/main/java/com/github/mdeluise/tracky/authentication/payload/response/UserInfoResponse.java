package com.github.mdeluise.tracky.authentication.payload.response;

import com.github.mdeluise.tracky.security.jwt.JwtTokenInfo;

public record UserInfoResponse(long id, String username, JwtTokenInfo jwt) {
}
