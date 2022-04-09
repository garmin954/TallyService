import React from "react";
import { View, Image, Text, ScrollView, ScrollViewProps } from "@tarojs/components";
import "./index.scss"

// ScrollViewProps
const Index = (props:ScrollViewProps & {children:JSX.Element[]}) => {

  props.refresherEnabled = true;
  props.refresherThreshold = 20
  props.refresherDefaultStyle = "red"
  props.refresherBackground = "blue"
  props.refresherTriggered= false
  props.pagingEnabled = true
  props.showScrollbar = false

  return (
    <ScrollView
      className="gm-scroll"
      scrollY
      scrollWithAnimation
      {...props}
    >
      {props.children}
    </ScrollView>
  );
};

export default Index;
