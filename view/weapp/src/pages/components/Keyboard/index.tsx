import { View } from "@tarojs/components";
import React, { useEffect, useState } from "react";
import style from "./index.module.scss"
const buttons = [
  {
    value: "1",
    gap: 'num'
  },
  {
    value: "2",
    gap: 'num'
  },
  {
    value: "3",
    gap: 'num'
  },
  {
    value: "del",
    gap: 'del'
  },
  {
    value: "4",
    gap: 'num'
  },
  {
    value: "5",
    gap: 'num'
  },
  {
    value: "6",
    gap: 'num'
  },
  {
    value: "确定",
    gap: 'button'
  },
  {
    value: "7",
    gap: 'num'
  },
  {
    value: "8",
    gap: 'num'
  },
  {
    value: "9",
    gap: 'num'
  },
  {
    value: "确定",
    gap: 'button'
  },
  {
    value: "0",
    gap: 'zero'
  },
  {
    value: "0",
    gap: 'zero'
  },
  {
    value: ".",
    gap: 'dot'
  },
  {
    value: "确定",
    gap: 'button'
  }
]

type Props = {
  value: string
  onChange: (val: string) => void
  onConfirm?: (val: string) => void
}

const Keyboard = (props: Props) => {
  const {value='', onChange, onConfirm} = props
  const [num, setNum] = useState<string>(value)

  useEffect(()=>{
    onChange(num)
  }, [num])

  const clickButton = (gap: string, value: string) => {
    switch (gap) {
      case "zero":
      case "num":
        const [before, after] = num.split(".")
        if(!num.includes(".") && before?.length >= 6) return
        if(after?.length >= 2) return
        setNum(num + value)
        break
      case "del":
        setNum(num.substring(0, num.length - 1))
        break
      case "dot":
        num.length && !num.includes(".") && setNum(num + value)
        break
      case "button":
        onConfirm && onConfirm(num)
      break
    }
  }

  return (<View className={style.keyboard}>
    {buttons.map(({ gap, value }) => <View onClick={() => clickButton(gap, value)} className={style[gap]}>{value}</View>)}
  </View>)
}

export default Keyboard
