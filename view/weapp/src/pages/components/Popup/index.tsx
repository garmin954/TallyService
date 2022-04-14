import React, { useEffect, useState } from "react";
import { ScrollView, View } from "@tarojs/components"
import style from "./index.module.scss"
import Taro, { Animation, nextTick, useReady } from "@tarojs/taro";
import IconFont from "@/utils/iconfont";

type Props = {
  hideHeader?: boolean // 显示隐藏header
  visible: boolean // 显示隐藏popup
  showMask?: boolean // 显示mask
  scrollY?: boolean // 上下滑动
  scrollX?: boolean // 左右滑动
  children ?: JSX.Element[] | JSX.Element // 插槽
  title?: string | JSX.Element // 标题
  onClose?: ()=>void // 关闭
}

const Popup = (props:Props) => {
  const {showMask=true,scrollY=true, scrollX= true, title='', visible= false, onClose, hideHeader=true} = props
  const [originVisible, setOriginVisible] = useState(visible)

  let timer
  useEffect(()=>{
    console.log("visible-------", visible)
    setOriginVisible(visible)
    // if(visible) {
    //   dh(true)
    //   return setOriginVisible(visible);
    // }

    // dh(false)
    // clearInterval(timer)
    // timer = setTimeout(()=>{
    //   setOriginVisible(visible)
    // }, 500)
  }, [visible])

  useReady(() => {
    // dh

    // setInterval(()=>{
      // dh(false)
    // }, 3000)
  })



  const [animationData, SetAnimationData] = useState<{actions: TaroGeneral.IAnyObject[];}>({actions:[]})

  var animation = Taro.createAnimation({
    duration: 500,
    timingFunction: "ease",
  })

  const dh = (start=true) =>{
    setTimeout(()=>{
      if(start){
        animation.translateZ(0).step()
      }else{
        animation.translateZ(1).step()
      }
      console.log("dhdhdhdhdhdhdhdhdh")
      SetAnimationData(animation.export())
    }, 100)
  }

  // if(!originVisible){
  //   return null
  // }

  return (
    <View animation={animationData} style={{transform: `translateZ(${originVisible?1:0})`}} className={style.drawer}>
      <View className={`${style.mask} ${showMask ?style.show:''}`} onClick={onClose}></View>
      <View  className={style.popup}>
        {hideHeader ? (
        <View className={style.header}>
          <View className={style.title}>{title}</View>
          <View className={style.close} onClick={onClose}>
            <IconFont color={"#488605"} size={50} name={"closefill"}></IconFont>
          </View>
        </View>
        ): null}
        <ScrollView
          scrollY={scrollY}
          scrollX={scrollX}
          className={style.container}
          >
          {props.children}
        </ScrollView>
      </View>
    </View>
  )
}

export default Popup
