/*
 Navicat Premium Data Transfer

 Source Server         : localhost_27017
 Source Server Type    : MongoDB
 Source Server Version : 40403
 Source Host           : localhost:27017
 Source Schema         : douyu

 Target Server Type    : MongoDB
 Target Server Version : 40403
 File Encoding         : 65001

 Date: 17/03/2021 12:23:24
*/


// ----------------------------
// Collection structure for send_gift
// ----------------------------
db.getCollection("send_gift").drop();
db.createCollection("send_gift");
db.getCollection("send_gift").createIndex({
    "to_uid": NumberInt("1")
}, {
    name: "idx_uid"
});
