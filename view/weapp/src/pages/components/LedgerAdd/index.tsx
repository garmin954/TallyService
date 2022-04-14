import React, { useState } from "react";
import {View, Text} from "@tarojs/components"
import { AtFab, AtFloatLayout } from "taro-ui";
import style from "./index.module.scss"
import { useReady } from "@tarojs/taro";
import IconFont from "@/utils/iconfont";
import Popup from "../Popup";

const LedgerAdd = () =>{

  useReady(()=>{
  })

  const [openForm, setOpenForm] = useState(false)
  return (
    <>
      <View className={style.addBtn}>
        <AtFab onClick={()=>setOpenForm(true)}>
          <IconFont color="#fff"  name="wirt" size={50} />
        </AtFab>
      </View>

      <Popup animationTime={600} visible={openForm} onClose={()=>setOpenForm(false)} title={'title'}>
        <View style={{height: '70vh'}}>Test</View>
      </Popup>
      <View>

      </View>
    </>
  )
}

export default LedgerAdd
