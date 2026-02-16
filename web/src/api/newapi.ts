import request from '@/utils/request'

// NewAPI 配置相关 API

export interface NewAPIProvider {
  name: string
  base_url: string
  api_key: string
  models: string[]
  priority: number
  weight: number
  enabled: boolean
}

export interface NewAPILoadBalancer {
  strategy: string
  providers: NewAPIProvider[]
}

export interface NewAPIConfig {
  base_url: string
  api_key: string
  load_balancer: NewAPILoadBalancer
}

export interface ProviderStats {
  name: string
  request_count: number
  error_count: number
  success_rate: number
  avg_latency: number
  last_used: string
}

/**
 * 获取NewAPI配置
 */
export function getNewAPIConfig() {
  return request({
    url: '/api/v1/newapi/config',
    method: 'get'
  })
}

/**
 * 更新NewAPI配置
 */
export function updateNewAPIConfig(data: NewAPIConfig) {
  return request({
    url: '/api/v1/newapi/config',
    method: 'put',
    data
  })
}

/**
 * 获取服务商统计
 */
export function getNewAPIStats() {
  return request({
    url: '/api/v1/newapi/stats',
    method: 'get'
  })
}

/**
 * 测试文本生成
 */
export function testTextGeneration(data: {
  model: string
  prompt: string
  system_prompt?: string
}) {
  return request({
    url: '/api/v1/newapi/text',
    method: 'post',
    data
  })
}

/**
 * 测试图像生成
 */
export function testImageGeneration(data: {
  model: string
  prompt: string
}) {
  return request({
    url: '/api/v1/newapi/image',
    method: 'post',
    data
  })
}
