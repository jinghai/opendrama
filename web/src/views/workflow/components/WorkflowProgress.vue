<template>
  <div class="workflow-progress">
    <el-steps :active="currentStep" finish-status="success" align-center>
      <el-step title="创建剧本" description="输入剧本基本信息和主题" />
      <el-step title="角色设计" description="AI自动提取或手动创建角色" />
      <el-step title="场景生成" description="生成剧中的场景背景" />
      <el-step title="分镜制作" description="AI生成详细分镜脚本" />
      <el-step title="图生视频" description="将图片转换为视频" />
      <el-step title="配音合成" description="TTS语音合成" />
      <el-step title="最终成片" description="合成完整短剧视频" />
    </el-steps>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

const props = defineProps<{
  step: number
}>()

const currentStep = computed(() => props.step || 0)
</script>

<style scoped lang="scss">
.workflow-progress {
  padding: 20px;
  background: var(--bg-card);
  border-radius: 8px;
  margin-bottom: 20px;
  
  :deep(.el-step__title) {
    font-size: 14px;
  }
  
  :deep(.el-step__description) {
    font-size: 12px;
  }
}
</style>
