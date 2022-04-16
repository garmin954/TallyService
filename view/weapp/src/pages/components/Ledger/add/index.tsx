import React, { useEffect, useState } from "react";
import { View, Text, Swiper, SwiperItem } from "@tarojs/components"
import { AtCalendar, AtFab, AtFloatLayout, AtTag } from "taro-ui";
import style from "./index.module.scss"
import Taro, { useReady } from "@tarojs/taro";
import IconFont from "@/utils/iconfont";
import Popup from "../../Popup";
import LedgerSelect from "../select";
import Keyboard from "../../Keyboard";
import { useRequest } from "taro-hooks";
import api from "@/api";

type StateForm = {
  ledgerIds: number[],
  amount: string
  memo: string
}
const PART_NUM: number = 12

const LedgerAdd = () => {
  const { data: classify, run: fetchClassify, loading } = useRequest(api.fetchClassify, { manual: true, formatResult: (res) => res.data, cacheKey: "classify" })
  useEffect(() => {
    formatClassify()
  }, [classify])

  const [state, setState] = useState<StateForm>({
    ledgerIds: [],
    amount: '',
    memo:''
  })

  const [classifyArr, setClassifyArr] = useState<ClassifyInfo[][]>([])
  // 格式化 分类tag
  const formatClassify = () => {
    if (classify?.length <= 0) return null
    // 分段
    const tagList: ClassifyInfo[][] = []
    for (const i in classify) {
      const k = Math.floor(+i / PART_NUM)
      if (!tagList?.[k]) tagList[k] = []
      tagList[k]?.push(classify[i])
    }
    setClassifyArr(tagList)
  }

  const [openForm, setOpenForm] = useState(false)
  useEffect(() => {
    openForm && fetchClassify({})
  }, [openForm])

  // 打开时间
  const [openTime, setOpenTime] = useState(false)

  // 图片上传
  const [openPic, setOpenPic] = useState(false)
  const chooseImage = () =>{
    Taro.chooseImage({
      success (res) {
        setOpenPic(true)
        const tempFilePaths = res.tempFilePaths
      }
    })
  }

  return (
    <>
      <View className={style.addBtn}>
        <AtFab onClick={() => setOpenForm(true)}>
          <IconFont color="#fff" name="wirt" size={50} />
        </AtFab>
      </View>

      <Popup
        animationTime={600}
        visible={openForm}
        onClose={() => setOpenForm(false)}
        title={<LedgerSelect
          value={state.ledgerIds}
          onChange={(ledgerIds) => {
            setState({ ...state, ledgerIds })
          }}
        />}
        catchMove={true}
      >
        <View className={style.container} catch-move={true}>
          <View className="flex jc-space-between ai-center">
            <View>
              <AtTag className="mr-20" active={true} type='primary'>支出</AtTag>
              <AtTag className="mr-20" active={true} type='primary'>入账</AtTag>
              <AtTag className="mr-20" active={true} type='primary'>不计入收支</AtTag>
            </View>
            <View onClick={() => setOpenTime(true)}>4月22日</View>
          </View>

          {/* 金额 */}
          <View className={style.amountBox}>
            <View className={style.unit}>￥</View>
            <View>{state.amount}<Text className={style.focus}>|</Text></View>
          </View>

          {/* 分类 */}
          <Swiper
            className={style.classify}
            indicatorColor='#999'
            indicatorActiveColor='#333'
            circular
          >
            {
              classifyArr?.map(item => (
                <SwiperItem className={style.swiperItem}>
                  {item.map(({ name, icon }) => (<View className={style.tagBox}>
                    <View className={style.icon}><IconFont size={60} name={icon}></IconFont></View>
                    <Text className={style.name}>{name}</Text>
                  </View>))}
                </SwiperItem>
              ))
            }
          </Swiper>

          {/* 备注，图片 */}
          <View className={style.memoPic}>
            <View className={style.memo}>
              {!state.memo ? '添加备注': state.memo}
            </View>
            <View className={style.date} onClick={() => setOpenTime(true)}>今天</View>
            <View className={style.tag}>#</View>
            <View className={style.pic} onClick={chooseImage}>pic</View>
          </View>
          {/* 键盘 */}
          <Keyboard
            value={`${state.amount}`}
            onChange={(amount) => {
              setState({ ...state, amount })
            }} />
        </View>
      </Popup>



      {/* 时间 */}
      <Popup
        animationTime={300}
        visible={openTime}
        onClose={() => setOpenTime(false)}
        catchMove={true}
      >
        <AtCalendar onDayClick={() => setOpenTime(false)} />
      </Popup>

      {/* 图片上传 */}
      <Popup
        animationTime={300}
        visible={openPic}
        onClose={() => setOpenPic(false)}
        catchMove={true}
      >
        <>图片</>
      </Popup>
    </>
  )
}

export default LedgerAdd
