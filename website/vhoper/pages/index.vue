<template>
  <no-ssr>
    <div>
      <a-carousel
        autoplay
        arrows
        effect="fade"
      >
        <div
          slot="prevArrow"
          class="custom-slick-arrow"
          style="left: 10px;zIndex: 1"
        >
          <a-icon type="left-circle" />
        </div>
        <div
          slot="nextArrow"
          class="custom-slick-arrow"
          style="right: 10px"
        >
          <a-icon type="right-circle" />
        </div>
        <div><h3>建设中，我觉得你看瞬间比较好</h3></div>
        <div><h3>建设中，我觉得你去聊天比较好</h3></div>
        <div><h3>建设中，我觉得你写博客比较好</h3></div>
        <div><h3>建设中，我觉得你比较好看</h3></div>
      </a-carousel>
    
      <a-row>
        <a-col :span="6">
          hoper
        </a-col>
        <a-col
          :span="12"
          style="text-align: center"
        >
          <div>
            原手机版作废，全部改成电脑版<br>
            先点上方的聊天来聊天吧
          </div>
        </a-col>
        <a-col :span="6">
          h
        </a-col>
      </a-row>
    </div>
  </no-ssr>
</template>

<script>
export default {}
</script>

<style scoped>
div .ant-carousel >>> div .slick-slide {
  text-align: center;
  height: 160px;
  line-height: 160px;
  background: #364d79;
  overflow: hidden;
}

div .ant-carousel >>> div .custom-slick-arrow {
  width: 25px;
  height: 25px;
  font-size: 25px;
  color: #fff;
  background-color: rgba(31, 45, 61, 0.11);
  opacity: 0.3;
}
div .ant-carousel >>> div .custom-slick-arrow:before {
  display: none;
}
div .ant-carousel >>> div .custom-slick-arrow:hover {
  opacity: 0.5;
}

div .ant-carousel >>> .slick-slide h3 {
  color: #fff;
}
</style>
