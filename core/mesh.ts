export interface MeshConfig {
  facilitatorUrl: string;
  merchantAddress: string;
  x402scanUrl?: string;
}

export interface ServiceRegistrationParams {
  name: string;
  description?: string;
  endpoint: string;
  category: string;
  price: string;
  network: 'base' | 'solana-mainnet';
  scheme: 'x402+eip712' | 'x402+solana';
  merchantAddress: string;
  metadata?: Record<string, any>;
}

export class AtlasMesh {
  private config: MeshConfig;
  private registeredServices: Map<string, ServiceRegistrationParams> = new Map();

  constructor(config: MeshConfig) {
    this.config = config;
  }

  async registerService(params: ServiceRegistrationParams): Promise<any> {
    const serviceId = this.generateServiceId();
    const priceMicro = Math.round(parseFloat(params.price) * 1_000_000).toString();

    const registrationData = {
      id: serviceId,
      name: params.name,
      description: params.description || '',
      endpoint: params.endpoint,
      category: params.category,
      network: params.network,
      merchantAddress: params.merchantAddress,
      accepts: [{
        asset: this.getAssetAddress(params.network),
        payTo: params.merchantAddress,
        network: params.network,
        maxAmountRequired: priceMicro,
        scheme: params.scheme,
        mimeType: 'application/json',
      }],
      metadata: params.metadata || {},
    };

    await this.registerWithFacilitator(registrationData);
    this.registeredServices.set(serviceId, params);

    return {
      serviceId,
      facilitatorUrl: `${this.config.facilitatorUrl}/discovery/resources/${serviceId}`,
    };
  }

  private generateServiceId(): string {
    return `service-${Date.now()}-${Math.random().toString(36).substr(2, 9)}`;
  }

  private getAssetAddress(network: 'base' | 'solana-mainnet'): string {
    if (network === 'base') {
      return '0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913';
    }
    return 'EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v';
  }

  private async registerWithFacilitator(data: any): Promise<void> {
    const response = await fetch(`${this.config.facilitatorUrl}/discovery/resources`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    });

    if (!response.ok) {
      throw new Error(`Registration failed: ${response.statusText}`);
    }
  }
}

export default AtlasMesh;

