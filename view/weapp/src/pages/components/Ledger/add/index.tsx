import React, { useCallback, useEffect, useState } from "react";
import { View, Text, Swiper, SwiperItem, Textarea } from "@tarojs/components"
import { AtCalendar, AtFab, AtTag } from "taro-ui";
import style from "./index.module.scss"
import Taro, { useReady } from "@tarojs/taro";
import IconFont from "@/utils/iconfont";
import Popup from "../../Popup";
import LedgerSelect from "../select";
import Keyboard from "../../Keyboard";
import { useRequest } from "taro-hooks";
import api from "@/api";

// 创建表单
type StateForm = {
  ledgerIds: number[],
  amount: string
  memo: string
  type: 1|2 // 1支出 2收入
}

// popup 状态
type OpenState = {
  form: boolean
  time: boolean
  pic: boolean
  memo: boolean
}

const RECORD_TYPE: Options[] = [
  {
    value: '1',
    label: '支出'
  },
  {
    value: '2',
    label: '入账'
  }
]

const PART_NUM: number = 12

const LedgerAdd = () => {
  const { data: classify, run: fetchClassify, loading } = useRequest(api.fetchClassify, { manual: true, formatResult: (res) => res.data, cacheKey: "classify" })
  useEffect(() => {
    formatClassify()
  }, [classify])

  // 表单值
  const [state, setState] = useState<StateForm>({
    ledgerIds: [],
    amount: '',
    memo: '',
    type: 1
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

  // 打开popup状态
  const [openState, setOpenState] = useState<OpenState>({
    form: false,
    pic: false,
    time: false,
    memo: false
  })

  useEffect(() => {
    openState.form && fetchClassify({})
  }, [openState.form])


  // 图片上传
  const chooseImage = () => {
    Taro.chooseImage({
      success(res) {
        setOpenState({ ...openState, pic: true })
        const tempFilePaths = res.tempFilePaths
      }
    })
  }


  return (
    <>
      <View className={style.addBtn}>
        <AtFab onClick={() => setOpenState({ ...openState, form: true })}>
          <IconFont color="#fff" name="wirt" size={50} />
        </AtFab>
      </View>

      <Popup
        animationTime={600}
        visible={openState.form}
        onClose={() => setOpenState({ ...openState, form: false })}
        title={<LedgerSelect
          value={state.ledgerIds}
          onChange={(ledgerIds) => {
            setState({ ...state, ledgerIds })
          }}
        />}
        catchMove={true}
      >
        <View className={style.container} catch-move={true}>
          <View className={'flex jc-space-between ai-center mb-40 '+style.pn30}>
            <View>
              {
                RECORD_TYPE.map(opt => <AtTag
                  className="mr-20"
                  active={+opt.value === state.type}
                  onClick={()=>{
                    setState({...state, type: +opt.value > 1 ? 2 : 1})
                  }}
                >
                  {opt.label}
                </AtTag>)
              }
            </View>
            <View onClick={() => setOpenState({ ...openState, time: true })}>4月22日</View>
          </View>

          {/* 金额 */}
          <View className={style.amountBox}>
            <View className={style.unit}>￥</View>
            <View className={style.amountTxt}>{state.amount}<Text className={style.focus}>|</Text></View>
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
            <View
              className={style.memo}
              onClick={() => setOpenState({ ...openState, memo: true })}
            >
              {!state.memo ? '添加备注...' : state.memo}
            </View>
            <View className={style.date} onClick={() => setOpenState({ ...openState, time: true })}>今天</View>
            <View className={style.tag}>#</View>
            <View className={style.pic} onClick={chooseImage}>pic</View>
          </View>
          {/* 键盘 */}
          <Keyboard
            className={style.pn30}
            value={`${state.amount}`}
            onChange={(amount) => {
              setState({ ...state, amount })
            }} />
        </View>
      </Popup>


      {/* 时间 */}
      <Popup
        animationTime={300}
        visible={openState.time}
        onClose={() => setOpenState({ ...openState, time: false })}
        catchMove={true}
      >
        <AtCalendar onDayClick={() => setOpenState({ ...openState, time: false })} />
      </Popup>

      {/* 图片上传 */}
      <Popup
        animationTime={300}
        visible={openState.pic}
        onClose={() => setOpenState({ ...openState, pic: false })}
        catchMove={true}
      >
        <>图片</>
      </Popup>

      {/* 备注 */}
      <Popup
        animationTime={350}
        visible={openState.memo}
        onClose={() => setOpenState({ ...openState, memo: false })}
        catchMove={true}
      >
        <View className={style.memoPopup}>
          <Textarea
          cursorSpacing={250}
          autoFocus={openState.memo}
          focus={openState.memo}
          />
        </View>
      </Popup>
    </>
  )
}

export default LedgerAdd
