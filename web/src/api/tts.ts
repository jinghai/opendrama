import request from '@/utils/request'

// TTS API接口

export interface TTSGenerateParams {
  provider: string
  voice: string
  text: string
  speed?: number
  pitch?: number
  format?: string
  save_to_file?: boolean
}

export interface TTSGenerateResponse {
  success: boolean
  audio_data?: string
  file_path?: string
  duration?: number
  provider: string
  voice: string
  latency: string
  error?: string
}

export interface TTSVoice {
  id: string
  name: string
  language: string
  gender: string
  style: string
  provider: string
}

export interface TTSBatchParams {
  provider: string
  voice: string
  texts: string[]
  speed?: number
  pitch?: number
  save_to_file?: boolean
}

export interface TTSBatchResponse {
  success: boolean
  results: TTSGenerateResponse[]
  total: number
  failed: number
}

/**
 * 生成TTS语音
 */
export function generateTTS(data: TTSGenerateParams) {
  return request({
    url: '/api/v1/tts/generate',
    method: 'post',
    data
  })
}

/**
 * 获取语音列表
 */
export function getTTSVoices(provider?: string) {
  return request({
    url: '/api/v1/tts/voices',
    method: 'get',
    params: { provider }
  })
}

/**
 * 获取TTS提供商列表
 */
export function getTTSProviders() {
  return request({
    url: '/api/v1/tts/providers',
    method: 'get'
  })
}

/**
 * 批量生成TTS语音
 */
export function batchGenerateTTS(data: TTSBatchParams) {
  return request({
    url: '/api/v1/tts/batch',
    method: 'post',
    data
  })
}
