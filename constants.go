package input

const (
	KeyReserved         Key = 0
	KeyEsc              Key = 1
	Key1                Key = 2
	Key2                Key = 3
	Key3                Key = 4
	Key4                Key = 5
	Key5                Key = 6
	Key6                Key = 7
	Key7                Key = 8
	Key8                Key = 9
	Key9                Key = 10
	Key0                Key = 11
	KeyMinus            Key = 12
	KeyEqual            Key = 13
	KeyBackspace        Key = 14
	KeyTab              Key = 15
	KeyQ                Key = 16
	KeyW                Key = 17
	KeyE                Key = 18
	KeyR                Key = 19
	KeyT                Key = 20
	KeyY                Key = 21
	KeyU                Key = 22
	KeyI                Key = 23
	KeyO                Key = 24
	KeyP                Key = 25
	KeyLeftbrace        Key = 26
	KeyRightbrace       Key = 27
	KeyEnter            Key = 28
	KeyLeftctrl         Key = 29
	KeyA                Key = 30
	KeyS                Key = 31
	KeyD                Key = 32
	KeyF                Key = 33
	KeyG                Key = 34
	KeyH                Key = 35
	KeyJ                Key = 36
	KeyK                Key = 37
	KeyL                Key = 38
	KeySemicolon        Key = 39
	KeyApostrophe       Key = 40
	KeyGrave            Key = 41
	KeyLeftshift        Key = 42
	KeyBackslash        Key = 43
	KeyZ                Key = 44
	KeyX                Key = 45
	KeyC                Key = 46
	KeyV                Key = 47
	KeyB                Key = 48
	KeyN                Key = 49
	KeyM                Key = 50
	KeyComma            Key = 51
	KeyDot              Key = 52
	KeySlash            Key = 53
	KeyRightshift       Key = 54
	KeyKpasterisk       Key = 55
	KeyLeftalt          Key = 56
	KeySpace            Key = 57
	KeyCapslock         Key = 58
	KeyF1               Key = 59
	KeyF2               Key = 60
	KeyF3               Key = 61
	KeyF4               Key = 62
	KeyF5               Key = 63
	KeyF6               Key = 64
	KeyF7               Key = 65
	KeyF8               Key = 66
	KeyF9               Key = 67
	KeyF10              Key = 68
	KeyNumlock          Key = 69
	KeyScrolllock       Key = 70
	KeyKp7              Key = 71
	KeyKp8              Key = 72
	KeyKp9              Key = 73
	KeyKpminus          Key = 74
	KeyKp4              Key = 75
	KeyKp5              Key = 76
	KeyKp6              Key = 77
	KeyKpplus           Key = 78
	KeyKp1              Key = 79
	KeyKp2              Key = 80
	KeyKp3              Key = 81
	KeyKp0              Key = 82
	KeyKpdot            Key = 83
	KeyZenkakuhankaku   Key = 85
	Key102nd            Key = 86
	KeyF11              Key = 87
	KeyF12              Key = 88
	KeyRo               Key = 89
	KeyKatakana         Key = 90
	KeyHiragana         Key = 91
	KeyHenkan           Key = 92
	KeyKatakanahiragana Key = 93
	KeyMuhenkan         Key = 94
	KeyKpjpcomma        Key = 95
	KeyKpenter          Key = 96
	KeyRightctrl        Key = 97
	KeyKpslash          Key = 98
	KeySysrq            Key = 99
	KeyRightalt         Key = 100
	KeyLinefeed         Key = 101
	KeyHome             Key = 102
	KeyUp               Key = 103
	KeyPageup           Key = 104
	KeyLeft             Key = 105
	KeyRight            Key = 106
	KeyEnd              Key = 107
	KeyDown             Key = 108
	KeyPagedown         Key = 109
	KeyInsert           Key = 110
	KeyDelete           Key = 111
	KeyMacro            Key = 112
	KeyMute             Key = 113
	KeyVolumedown       Key = 114
	KeyVolumeup         Key = 115
	KeyPower            Key = 116
	KeyKpequal          Key = 117
	KeyKpplusminus      Key = 118
	KeyPause            Key = 119
	KeyScale            Key = 120
	KeyKpcomma          Key = 121
	KeyHangeul          Key = 122
	KeyHanguel          Key = KeyHangeul
	KeyHanja            Key = 123
	KeyYen              Key = 124
	KeyLeftmeta         Key = 125
	KeyRightmeta        Key = 126
	KeyCompose          Key = 127
	KeyStop             Key = 128
	KeyAgain            Key = 129
	KeyProps            Key = 130
	KeyUndo             Key = 131
	KeyFront            Key = 132
	KeyCopy             Key = 133
	KeyOpen             Key = 134
	KeyPaste            Key = 135
	KeyFind             Key = 136
	KeyCut              Key = 137
	KeyHelp             Key = 138
	KeyMenu             Key = 139
	KeyCalc             Key = 140
	KeySetup            Key = 141
	KeySleep            Key = 142
	KeyWakeup           Key = 143
	KeyFile             Key = 144
	KeySendfile         Key = 145
	KeyDeletefile       Key = 146
	KeyXfer             Key = 147
	KeyProg1            Key = 148
	KeyProg2            Key = 149
	KeyWww              Key = 150
	KeyMsdos            Key = 151
	KeyCoffee           Key = 152
	KeyScreenlock       Key = KeyCoffee
	KeyDirection        Key = 153
	KeyCyclewindows     Key = 154
	KeyMail             Key = 155
	KeyBookmarks        Key = 156
	KeyComputer         Key = 157
	KeyBack             Key = 158
	KeyForward          Key = 159
	KeyClosecd          Key = 160
	KeyEjectcd          Key = 161
	KeyEjectclosecd     Key = 162
	KeyNextsong         Key = 163
	KeyPlaypause        Key = 164
	KeyPrevioussong     Key = 165
	KeyStopcd           Key = 166
	KeyRecord           Key = 167
	KeyRewind           Key = 168
	KeyPhone            Key = 169
	KeyIso              Key = 170
	KeyConfig           Key = 171
	KeyHomepage         Key = 172
	KeyRefresh          Key = 173
	KeyExit             Key = 174
	KeyMove             Key = 175
	KeyEdit             Key = 176
	KeyScrollup         Key = 177
	KeyScrolldown       Key = 178
	KeyKpleftparen      Key = 179
	KeyKprightparen     Key = 180
	KeyNew              Key = 181
	KeyRedo             Key = 182
	KeyF13              Key = 183
	KeyF14              Key = 184
	KeyF15              Key = 185
	KeyF16              Key = 186
	KeyF17              Key = 187
	KeyF18              Key = 188
	KeyF19              Key = 189
	KeyF20              Key = 190
	KeyF21              Key = 191
	KeyF22              Key = 192
	KeyF23              Key = 193
	KeyF24              Key = 194
	KeyPlaycd           Key = 200
	KeyPausecd          Key = 201
	KeyProg3            Key = 202
	KeyProg4            Key = 203
	KeyDashboard        Key = 204
	KeySuspend          Key = 205
	KeyClose            Key = 206
	KeyPlay             Key = 207
	KeyFastforward      Key = 208
	KeyBassboost        Key = 209
	KeyPrint            Key = 210
	KeyHp               Key = 211
	KeyCamera           Key = 212
	KeySound            Key = 213
	KeyQuestion         Key = 214
	KeyEmail            Key = 215
	KeyChat             Key = 216
	KeySearch           Key = 217
	KeyConnect          Key = 218
	KeyFinance          Key = 219
	KeySport            Key = 220
	KeyShop             Key = 221
	KeyAlterase         Key = 222
	KeyCancel           Key = 223
	KeyBrightnessdown   Key = 224
	KeyBrightnessup     Key = 225
	KeyMedia            Key = 226
	KeySwitchvideomode  Key = 227
	KeyKbdillumtoggle   Key = 228
	KeyKbdillumdown     Key = 229
	KeyKbdillumup       Key = 230
	KeySend             Key = 231
	KeyReply            Key = 232
	KeyForwardmail      Key = 233
	KeySave             Key = 234
	KeyDocuments        Key = 235
	KeyBattery          Key = 236
	KeyBluetooth        Key = 237
	KeyWlan             Key = 238
	KeyUwb              Key = 239
	KeyUnknown          Key = 240
	KeyVideoNext        Key = 241
	KeyVideoPrev        Key = 242
	KeyBrightnessCycle  Key = 243
	KeyBrightnessZero   Key = 244
	KeyDisplayOff       Key = 245
	KeyWimax            Key = 246
	KeyRfkill           Key = 247
	KeyMicmute          Key = 248
	BtnMisc             Key = 0x100
	Btn0                Key = 0x100
	Btn1                Key = 0x101
	Btn2                Key = 0x102
	Btn3                Key = 0x103
	Btn4                Key = 0x104
	Btn5                Key = 0x105
	Btn6                Key = 0x106
	Btn7                Key = 0x107
	Btn8                Key = 0x108
	Btn9                Key = 0x109
	BtnMouse            Key = 0x110
	BtnLeft             Key = 0x110
	BtnRight            Key = 0x111
	BtnMiddle           Key = 0x112
	BtnSide             Key = 0x113
	BtnExtra            Key = 0x114
	BtnForward          Key = 0x115
	BtnBack             Key = 0x116
	BtnTask             Key = 0x117
	BtnJoystick         Key = 0x120
	BtnTrigger          Key = 0x120
	BtnThumb            Key = 0x121
	BtnThumb2           Key = 0x122
	BtnTop              Key = 0x123
	BtnTop2             Key = 0x124
	BtnPinkie           Key = 0x125
	BtnBase             Key = 0x126
	BtnBase2            Key = 0x127
	BtnBase3            Key = 0x128
	BtnBase4            Key = 0x129
	BtnBase5            Key = 0x12a
	BtnBase6            Key = 0x12b
	BtnDead             Key = 0x12f
	BtnGamepad          Key = 0x130
	BtnA                Key = 0x130
	BtnB                Key = 0x131
	BtnC                Key = 0x132
	BtnX                Key = 0x133
	BtnY                Key = 0x134
	BtnZ                Key = 0x135
	BtnTl               Key = 0x136
	BtnTr               Key = 0x137
	BtnTl2              Key = 0x138
	BtnTr2              Key = 0x139
	BtnSelect           Key = 0x13a
	BtnStart            Key = 0x13b
	BtnMode             Key = 0x13c
	BtnThumbl           Key = 0x13d
	BtnThumbr           Key = 0x13e
	BtnDigi             Key = 0x140
	BtnToolPen          Key = 0x140
	BtnToolRubber       Key = 0x141
	BtnToolBrush        Key = 0x142
	BtnToolPencil       Key = 0x143
	BtnToolAirbrush     Key = 0x144
	BtnToolFinger       Key = 0x145
	BtnToolMouse        Key = 0x146
	BtnToolLens         Key = 0x147
	BtnToolQuinttap     Key = 0x148
	BtnTouch            Key = 0x14a
	BtnStylus           Key = 0x14b
	BtnStylus2          Key = 0x14c
	BtnToolDoubletap    Key = 0x14d
	BtnToolTripletap    Key = 0x14e
	BtnToolQuadtap      Key = 0x14f
	BtnWheel            Key = 0x150
	BtnGearDown         Key = 0x150
	BtnGearUp           Key = 0x151
	KeyOk               Key = 0x160
	KeySelect           Key = 0x161
	KeyGoto             Key = 0x162
	KeyClear            Key = 0x163
	KeyPower2           Key = 0x164
	KeyOption           Key = 0x165
	KeyInfo             Key = 0x166
	KeyTime             Key = 0x167
	KeyVendor           Key = 0x168
	KeyArchive          Key = 0x169
	KeyProgram          Key = 0x16a
	KeyChannel          Key = 0x16b
	KeyFavorites        Key = 0x16c
	KeyEpg              Key = 0x16d
	KeyPvr              Key = 0x16e
	KeyMhp              Key = 0x16f
	KeyLanguage         Key = 0x170
	KeyTitle            Key = 0x171
	KeySubtitle         Key = 0x172
	KeyAngle            Key = 0x173
	KeyZoom             Key = 0x174
	KeyMode             Key = 0x175
	KeyKeyboard         Key = 0x176
	KeyScreen           Key = 0x177
	KeyPc               Key = 0x178
	KeyTv               Key = 0x179
	KeyTv2              Key = 0x17a
	KeyVcr              Key = 0x17b
	KeyVcr2             Key = 0x17c
	KeySat              Key = 0x17d
	KeySat2             Key = 0x17e
	KeyCd               Key = 0x17f
	KeyTape             Key = 0x180
	KeyRadio            Key = 0x181
	KeyTuner            Key = 0x182
	KeyPlayer           Key = 0x183
	KeyText             Key = 0x184
	KeyDvd              Key = 0x185
	KeyAux              Key = 0x186
	KeyMp3              Key = 0x187
	KeyAudio            Key = 0x188
	KeyVideo            Key = 0x189
	KeyDirectory        Key = 0x18a
	KeyList             Key = 0x18b
	KeyMemo             Key = 0x18c
	KeyCalendar         Key = 0x18d
	KeyRed              Key = 0x18e
	KeyGreen            Key = 0x18f
	KeyYellow           Key = 0x190
	KeyBlue             Key = 0x191
	KeyChannelup        Key = 0x192
	KeyChanneldown      Key = 0x193
	KeyFirst            Key = 0x194
	KeyLast             Key = 0x195
	KeyAb               Key = 0x196
	KeyNext             Key = 0x197
	KeyRestart          Key = 0x198
	KeySlow             Key = 0x199
	KeyShuffle          Key = 0x19a
	KeyBreak            Key = 0x19b
	KeyPrevious         Key = 0x19c
	KeyDigits           Key = 0x19d
	KeyTeen             Key = 0x19e
	KeyTwen             Key = 0x19f
	KeyVideophone       Key = 0x1a0
	KeyGames            Key = 0x1a1
	KeyZoomin           Key = 0x1a2
	KeyZoomout          Key = 0x1a3
	KeyZoomreset        Key = 0x1a4
	KeyWordprocessor    Key = 0x1a5
	KeyEditor           Key = 0x1a6
	KeySpreadsheet      Key = 0x1a7
	KeyGraphicseditor   Key = 0x1a8
	KeyPresentation     Key = 0x1a9
	KeyDatabase         Key = 0x1aa
	KeyNews             Key = 0x1ab
	KeyVoicemail        Key = 0x1ac
	KeyAddressbook      Key = 0x1ad
	KeyMessenger        Key = 0x1ae
	KeyDisplaytoggle    Key = 0x1af
	KeySpellcheck       Key = 0x1b0
	KeyLogoff           Key = 0x1b1
	KeyDollar           Key = 0x1b2
	KeyEuro             Key = 0x1b3
	KeyFrameback        Key = 0x1b4
	KeyFrameforward     Key = 0x1b5
	KeyContextMenu      Key = 0x1b6
	KeyMediaRepeat      Key = 0x1b7
	Key10channelsup     Key = 0x1b8
	Key10channelsdown   Key = 0x1b9
	KeyImages           Key = 0x1ba
	KeyDelEol           Key = 0x1c0
	KeyDelEos           Key = 0x1c1
	KeyInsLine          Key = 0x1c2
	KeyDelLine          Key = 0x1c3
	KeyFn               Key = 0x1d0
	KeyFnEsc            Key = 0x1d1
	KeyFnF1             Key = 0x1d2
	KeyFnF2             Key = 0x1d3
	KeyFnF3             Key = 0x1d4
	KeyFnF4             Key = 0x1d5
	KeyFnF5             Key = 0x1d6
	KeyFnF6             Key = 0x1d7
	KeyFnF7             Key = 0x1d8
	KeyFnF8             Key = 0x1d9
	KeyFnF9             Key = 0x1da
	KeyFnF10            Key = 0x1db
	KeyFnF11            Key = 0x1dc
	KeyFnF12            Key = 0x1dd
	KeyFn1              Key = 0x1de
	KeyFn2              Key = 0x1df
	KeyFnD              Key = 0x1e0
	KeyFnE              Key = 0x1e1
	KeyFnF              Key = 0x1e2
	KeyFnS              Key = 0x1e3
	KeyFnB              Key = 0x1e4
	KeyBrlDot1          Key = 0x1f1
	KeyBrlDot2          Key = 0x1f2
	KeyBrlDot3          Key = 0x1f3
	KeyBrlDot4          Key = 0x1f4
	KeyBrlDot5          Key = 0x1f5
	KeyBrlDot6          Key = 0x1f6
	KeyBrlDot7          Key = 0x1f7
	KeyBrlDot8          Key = 0x1f8
	KeyBrlDot9          Key = 0x1f9
	KeyBrlDot10         Key = 0x1fa
	KeyNumeric0         Key = 0x200
	KeyNumeric1         Key = 0x201
	KeyNumeric2         Key = 0x202
	KeyNumeric3         Key = 0x203
	KeyNumeric4         Key = 0x204
	KeyNumeric5         Key = 0x205
	KeyNumeric6         Key = 0x206
	KeyNumeric7         Key = 0x207
	KeyNumeric8         Key = 0x208
	KeyNumeric9         Key = 0x209
	KeyNumericStar      Key = 0x20a
	KeyNumericPound     Key = 0x20b
	KeyCameraFocus      Key = 0x210
	KeyWpsButton        Key = 0x211
	KeyTouchpadToggle   Key = 0x212
	KeyTouchpadOn       Key = 0x213
	KeyTouchpadOff      Key = 0x214
	KeyCameraZoomin     Key = 0x215
	KeyCameraZoomout    Key = 0x216
	KeyCameraUp         Key = 0x217
	KeyCameraDown       Key = 0x218
	KeyCameraLeft       Key = 0x219
	KeyCameraRight      Key = 0x21a
	BtnTriggerHappy     Key = 0x2c0
	BtnTriggerHappy1    Key = 0x2c0
	BtnTriggerHappy2    Key = 0x2c1
	BtnTriggerHappy3    Key = 0x2c2
	BtnTriggerHappy4    Key = 0x2c3
	BtnTriggerHappy5    Key = 0x2c4
	BtnTriggerHappy6    Key = 0x2c5
	BtnTriggerHappy7    Key = 0x2c6
	BtnTriggerHappy8    Key = 0x2c7
	BtnTriggerHappy9    Key = 0x2c8
	BtnTriggerHappy10   Key = 0x2c9
	BtnTriggerHappy11   Key = 0x2ca
	BtnTriggerHappy12   Key = 0x2cb
	BtnTriggerHappy13   Key = 0x2cc
	BtnTriggerHappy14   Key = 0x2cd
	BtnTriggerHappy15   Key = 0x2ce
	BtnTriggerHappy16   Key = 0x2cf
	BtnTriggerHappy17   Key = 0x2d0
	BtnTriggerHappy18   Key = 0x2d1
	BtnTriggerHappy19   Key = 0x2d2
	BtnTriggerHappy20   Key = 0x2d3
	BtnTriggerHappy21   Key = 0x2d4
	BtnTriggerHappy22   Key = 0x2d5
	BtnTriggerHappy23   Key = 0x2d6
	BtnTriggerHappy24   Key = 0x2d7
	BtnTriggerHappy25   Key = 0x2d8
	BtnTriggerHappy26   Key = 0x2d9
	BtnTriggerHappy27   Key = 0x2da
	BtnTriggerHappy28   Key = 0x2db
	BtnTriggerHappy29   Key = 0x2dc
	BtnTriggerHappy30   Key = 0x2dd
	BtnTriggerHappy31   Key = 0x2de
	BtnTriggerHappy32   Key = 0x2df
	BtnTriggerHappy33   Key = 0x2e0
	BtnTriggerHappy34   Key = 0x2e1
	BtnTriggerHappy35   Key = 0x2e2
	BtnTriggerHappy36   Key = 0x2e3
	BtnTriggerHappy37   Key = 0x2e4
	BtnTriggerHappy38   Key = 0x2e5
	BtnTriggerHappy39   Key = 0x2e6
	BtnTriggerHappy40   Key = 0x2e7
	KeyMinInteresting   Key = KeyMute
	KeyMax              Key = 0x2ff
)
