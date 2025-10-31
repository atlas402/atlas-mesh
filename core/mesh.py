from typing import Optional, Dict, Any
from dataclasses import dataclass
import uuid
import aiohttp

@dataclass
class ServiceRegistrationParams:
    name: str
    description: Optional[str]
    endpoint: str
    category: str
    price: str
    network: str
    scheme: str
    merchant_address: str
    metadata: Optional[Dict[str, Any]] = None

class AtlasMesh:
    def __init__(
        self,
        facilitator_url: str,
        merchant_address: str,
        x402scan_url: Optional[str] = None
    ):
        self.facilitator_url = facilitator_url
        self.merchant_address = merchant_address
        self.x402scan_url = x402scan_url
        self.services: Dict[str, ServiceRegistrationParams] = {}
        
    async def register_service(
        self,
        params: ServiceRegistrationParams
    ) -> Dict[str, Any]:
        service_id = str(uuid.uuid4())
        price_micro = str(int(float(params.price) * 1_000_000))
        
        registration_data = {
            'id': service_id,
            'name': params.name,
            'description': params.description or '',
            'endpoint': params.endpoint,
            'category': params.category,
            'network': params.network,
            'merchant_address': params.merchant_address,
            'accepts': [{
                'asset': self._get_asset_address(params.network),
                'payTo': params.merchant_address,
                'network': params.network,
                'maxAmountRequired': price_micro,
                'scheme': params.scheme,
                'mimeType': 'application/json',
            }],
            'metadata': params.metadata or {},
        }
        
        await self._register_with_facilitator(registration_data)
        
        self.services[service_id] = params
        
        return {
            'service_id': service_id,
            'facilitator_url': f"{self.facilitator_url}/discovery/resources/{service_id}",
        }
    
    def _get_asset_address(self, network: str) -> str:
        if network == 'base':
            return '0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913'
        else:
            return 'EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v'
    
    async def _register_with_facilitator(self, data: Dict[str, Any]):
        async with aiohttp.ClientSession() as session:
            async with session.post(
                f"{self.facilitator_url}/discovery/resources",
                json=data
            ) as response:
                response.raise_for_status()



