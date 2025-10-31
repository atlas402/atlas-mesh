package com.atlas402.mesh.core;

import java.util.concurrent.CompletableFuture;

public class AtlasMesh {
    private final String facilitatorUrl;
    private final String merchantAddress;
    
    public AtlasMesh(String facilitatorUrl, String merchantAddress) {
        this.facilitatorUrl = facilitatorUrl;
        this.merchantAddress = merchantAddress;
    }
    
    public CompletableFuture<ServiceRegistrationResult> registerService(ServiceRegistrationParams params) {
        return CompletableFuture.supplyAsync(() -> {
            String serviceId = generateServiceId();
            return new ServiceRegistrationResult(
                serviceId,
                facilitatorUrl + "/discovery/resources/" + serviceId
            );
        });
    }
    
    private String generateServiceId() {
        return "service-" + System.currentTimeMillis();
    }
}



