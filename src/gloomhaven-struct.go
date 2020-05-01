package main

// using information from https://kb.tabletopsimulator.com/custom-content/save-file-format/
// maybe use for custom decks, states, etc: https://stackoverflow.com/questions/21268000/unmarshaling-nested-json-objects

/*  SaveState //Holds a state of the game
    public string SaveName = "";
    public string GameMode = "";
    public float Gravity = 0.5f;
    public float PlayArea = 0.5f;
    public string Date = "";
    public string Table = "";
    [Tag(TagType.URL)]
    public string TableURL = null; //For custom tables
    public string Sky = "";
    [Tag(TagType.URL)]
    public string SkyURL = null; //For custom sky
    public string Note = "";
    public string Rules = "";
    public string XmlUI = ""; //Custom Xml UI
    public List<CustomAssetState> CustomUIAssets;
    public string LuaScript = "";
    public string LuaScriptState = ""; // Serialized running Lua code
    public GridState Grid; //Grid menu settings
    public LightingState Lighting; //Lighting menu settings
    public HandsState Hands; //Hand menu settings and hand positions
    public TurnsState Turns; //Turn menu settings
    public List<VectorLineState> VectorLines; // Vector lines on canvas 0 (table + beyond)
    public List<ObjectState> ObjectStates; //Objects on the table
    public List<SnapPointState> SnapPoints; //Snap points not attached to objects
    public List<CustomDecalState> DecalPallet; //Decals that can be placed in the world
    public List<DecalState> Decals; //Decals not attached to objects
    public Dictionary<int, TabState> TabStates = new Dictionary<int, TabState>(); //Notepad tabs
    public CameraState[] CameraStates; //Saved camera positions
    public string VersionNumber = "";*/
type SaveState struct {
	SaveName       string        `json:"SaveName"`
	GameMode       string        `json:"GameMode"`
	Gravity        float64       `json:"Gravity"`
	PlayArea       float64       `json:"PlayArea"`
	Date           string        `json:"Date"`
	Table          string        `json:"Table"`
	Sky            string        `json:"Sky"`
	SkyURL         string        `json:"SkyURL"`
	Note           string        `json:"Note"`
	Rules          string        `json:"Rules"`
	XMLUI          string        `json:"XmlUI"`
	CustomUIAssets []CustomAsset `json:"CustomUIAssets"`
	LuaScript      string        `json:"LuaScript"`
	LuaScriptState string        `json:"LuaScriptState"`
	MusicPlayer    MusicPlayer   `json:"MusicPlayer"`
	Grid           Grid          `json:"Grid"`
	Lighting       Lighting      `json:"Lighting"`
	Hands          Hands         `json:"Hands"`
	Turns          Turns         `json:"Turns"`
	ObjectStates   []ObjectState `json:"ObjectStates"`
	SnapPoints     []SnapPoint   `json:"SnapPoints"`
	DecalPallets   []DecalPallet `json:"DecalPallet"` //TODO: Not sure if this is right
	TabStates      []Tab         `json:"TabStates"`
	CameraStates   []CameraState `json:"CameraStates"`
	VersionNumber  string        `json:"VersionNumber"`
}

/* ObjectState //Moveable objects
	public string Name = ""; //Internal object name
    public TransformState Transform; //Position, Rotation, Scale
    public string Nickname = ""; //Name supplied in game
    public string Description = "";
    public ColourState ColorDiffuse; //Material color tint
    // Toggles
    public bool Locked = false; //Freeze object in place
    public bool Grid = true; //Object will snap to grid
    public bool Snap = true; //Object will snap to snap points
    public bool Autoraise = true; //Object will raise above others and avoid collision
    public bool Sticky = true; //When picked up objects above this one will be attached to it
    public bool Tooltip = true; //When hovering object will display tooltips
    public bool GridProjection = false; //Grid will project on this object
    // Nullable to hide object specific properties and save space
    public bool? HideWhenFaceDown; //When face down object is question mark hidden
    public bool? Hands; //Object will enter player hands
    public bool? AltSound; //Some objects have 2 materials, with two sound sets
    public int? MaterialIndex; //Some objects can have multiple materials
    public int? MeshIndex; //Some objects can have multiple meshes
    public int? Layer; //Sound Layer
    public int? Number;
    public int? CardID;
    public bool? SidewaysCard;
    public bool? RPGmode;
    public bool? RPGdead;
    public string FogColor = null;
    public bool? FogHidePointers;
    public bool? FogReverseHiding;
    public bool? FogSeethrough;
    public List<int> DeckIDs;
    public Dictionary<int, CustomDeckState> CustomDeck; //Key matches the hundreth place of the id (ex. id = 354, index = 3)
    public CustomMeshState CustomMesh;
    public CustomImageState CustomImage;
    public CustomAssetbundleState CustomAssetbundle;
    public FogOfWarSaveState FogOfWar;
    public FogOfWarRevealerSaveState FogOfWarRevealer;
    public ClockSaveState Clock;
    public CounterState Counter;
    public TabletState Tablet;
    public Mp3PlayerState Mp3Player;
    public CalculatorState Calculator;
    public TextState Text;
    public string XmlUI = ""; //Custom Xml UI
    public List<CustomAssetState> CustomUIAssets;
    public string LuaScript = "";
    public string LuaScriptState = ""; // Serialized running Lua code
    public List<ObjectState> ContainedObjects; //Objects inside this one
    public PhysicsMaterialState PhysicsMaterial; //Use to modify the physics material (friction, bounce, etc.) http://docs.unity3d.com/Manual/class-PhysicMaterial.html
    public RigidbodyState Rigidbody; //Use to modify the physical properties (mass, drag, etc) http://docs.unity3d.com/Manual/class-Rigidbody.html
    public JointFixedState JointFixed; //Joints can be used to attached/link objects together check the classes below
    public JointHingeState JointHinge;
    public JointSpringState JointSpring;
    public string GUID = ""; //Used so objects can reference other objects, ex. joints or scripting
    public List<SnapPointState> AttachedSnapPoints; //Snap points that are stuck to this object, happens when placing a snap point on a locked object
    public List<VectorLineState> AttachedVectorLines; // Vector lines that are stuck to this object, happens when drawing a vector line on a locked object
    public List<DecalState> AttachedDecals; //Decals that are attached to this objects
    public Dictionary<int, ObjectState> States; //Objects can have multiple states which can be swapped between
    public List<RotationValueState> RotationValues; //Rotation values are tooltip values tied to rotations */
type ObjectState struct {
	Name         string    `json:"Name"` //Internal object name
	Transform    Transform `json:"Transform"`
	Nickname     string    `json:"Nickname"` //Name supplied in game
	Description  string    `json:"Description"`
	GMNotes      string    `json:"GMNotes"`
	ColorDiffuse Color     `json:"ColorDiffuse,omitempty"`
	// Toggles
	Locked         bool `json:"Locked"`
	Grid           bool `json:"Grid"`
	Snap           bool `json:"Snap"`
	IgnoreFoW      bool `json:"IgnoreFoW"`
	Autoraise      bool `json:"Autoraise"`
	Sticky         bool `json:"Sticky"`
	Tooltip        bool `json:"Tooltip"`
	GridProjection bool `json:"GridProjection"`
	// Nullable to hide object specific properties and save space
	HideWhenFaceDown    bool              `json:"HideWhenFaceDown,omitempty"`
	Hands               bool              `json:"Hands,omitempty"`
	AltSound            bool              `json:"AltSound,omitempty"`
	MaterialIndex       int               `json:"MaterialIndex,omitempty"`
	MeshIndex           int               `json:"MeshIndex,omitempty"`
	Layer               int               `json:"Layer,omitempty"`
	Number              int               `json:"Number,omitempty"`
	CardID              int               `json:"CardID,omitempty"`
	SidewaysCard        bool              `json:"SidewaysCard,omitempty"`
	RPGmode             bool              `json:"RPGmode,omitempty"`
	RPGdead             bool              `json:"RPGdead,omitempty"`
	FogColor            string            `json:"FogColor,omitempty"`
	FogHidePointers     bool              `json:"FogHidePointers,omitempty"`
	FogReverseHiding    bool              `json:"FogReverseHiding,omitempty"`
	FogSeethrough       bool              `json:"FogSeethrough,omitempty"`
	DeckIDs             []int             `json:"DeckIDs,omitempty"`
	CustomDecks         []CustomDeck      `json:"CustomDeck,omitempty"`
	CustomMesh          CustomMesh        `json:"CustomMesh,omitempty"`
	CustomImage         CustomImage       `json:"CustomImage,omitempty"`
	CustomAssetbundle   CustomAssetbundle `json:"CustomAssetbundle,omitempty"`
	FogOfWar            FogOfWar          `json:"FogOfWar,omitempty"`
	FogOfWarRevealer    FogOfWarRevealer  `json:"FogOfWarRevealer,omitempty"`
	Clock               Clock             `json:"Clock,omitempty"`
	Counter             Counter           `json:"Counter,omitempty"`
	Tablet              Tablet            `json:"Tablet,omitempty"`
	Mp3Player           Mp3Player         `json:"Mp3Player,omitempty"`
	Calculator          Calculator        `json:"Calculator,omitempty"`
	Text                Text              `json:"Text,omitempty"`
	XMLUI               string            `json:"XmlUI,omitempty"`
	CustomUIAssets      []CustomAsset     `json:"CustomUIAssets,omitempty"`
	LuaScript           string            `json:"LuaScript,omitempty"`
	LuaScriptState      string            `json:"LuaScriptState,omitempty"`
	ContainedObjects    []ObjectState     `json:"ContainedObjects,omitempty"`
	PhysicsMaterial     PhysicsMaterial   `json:"PhysicsMaterial,omitempty"`
	Rigidbody           Rigidbody         `json:"Rigidbody,omitempty"`
	JointFixed          JointFixed        `json:"JointFixed,omitempty"`
	JointHinge          JointHinge        `json:"JointHinge,omitempty"`
	JointSpring         JointSpring       `json:"JointSpring,omitempty"`
	GUID                string            `json:"GUID"`
	AttachedSnapPoints  []SnapPoint       `json:"AttachedSnapPoints,omitempty"`
	AttachedVectorLines []VectorLine      `json:"AttachedVectorLines,omitempty"`
	AttachedDecals      []Decal           `json:"AttachedDecals,omitempty"`
	States              []ObjectState     `json:"States,omitempty"`
	RotationValues      []Vector          `json:"RotationValues,omitempty"`
	CustomPDF           CustomPDF         `json:"CustomPDF,omitempty"`
}

/* GridState
	public GridType Type = GridType.Box;
    public bool Lines = false;
    public ColourState Color = new ColourState(0,0,0);
    public float Opacity = 0.75f; //0-1 Alpha opacity
    public bool ThickLines = false;
    public bool Snapping = false; //Line snapping
    public bool Offset = false; //Center snapping
    public bool BothSnapping = false; //Both snapping
    public float xSize = 2f;
    public float ySize = 2f;
    public VectorState PosOffset = new VectorState(0,1,0);*/
type Grid struct {
	Type         int     `json:"Type"` //Box[0], HexHorizontal[1], HexVertical[2]
	Lines        bool    `json:"Lines"`
	Color        Color   `json:"Color"`
	Opacity      float64 `json:"Opacity"`
	ThickLines   bool    `json:"ThickLines"`
	Snapping     bool    `json:"Snapping"`
	Offset       bool    `json:"Offset"`
	BothSnapping bool    `json:"BothSnapping"`
	XSize        float64 `json:"xSize"`
	YSize        float64 `json:"ySize"`
	PosOffset    Vector  `json:"PosOffset"`
}

/* LightingState
   public float LightIntensity = 0.54f; //0-8
   public ColourState LightColor = new ColourState(1f, 0.9804f, 0.8902f);
   public float AmbientIntensity = 1.3f; //0-8
   public AmbientType AmbientType = AmbientType.Background;
   public ColourState AmbientSkyColor = new ColourState(0.5f, 0.5f, 0.5f);
   public ColourState AmbientEquatorColor = new ColourState(0.5f, 0.5f, 0.5f);
   public ColourState AmbientGroundColor = new ColourState(0.5f, 0.5f, 0.5f);
   public float ReflectionIntensity = 1f; //0-1
   public int LutIndex = 0;
   public float LutContribution = 1f; //0-1
   [Tag(TagType.URL)]
   public string LutURL; //LUT 256x16*/
type Lighting struct {
	LightIntensity      float64 `json:"LightIntensity"`
	LightColor          Color   `json:"LightColor"`
	AmbientIntensity    float64 `json:"AmbientIntensity"`
	AmbientType         int     `json:"AmbientType"` //Background[0] = ambient light comes from the background, Gradient[1] = ambient light comes from the three ambient colors
	AmbientSkyColor     Color   `json:"AmbientSkyColor"`
	AmbientEquatorColor Color   `json:"AmbientEquatorColor"`
	AmbientGroundColor  Color   `json:"AmbientGroundColor"`
	ReflectionIntensity float64 `json:"ReflectionIntensity"`
	LutIndex            int     `json:"LutIndex"`
	LutContribution     float64 `json:"LutContribution"`
	LutURL              string  `json:"LutURL"`
}

/* Hands
   public bool Enable = true;
   public bool DisableUnused = false;
   public HidingType Hiding = HidingType.Default;
   public List<HandTransformState> HandTransforms;*/
type Hands struct {
	Enable         bool            `json:"Enable"`
	DisableUnused  bool            `json:"DisableUnused"`
	Hiding         int             `json:"Hiding"` //Default[0] = only owner can see, Reverse[1] = opposite of default, Disable[2] = hiding is disabled
	HandTransforms []HandTransform `json:"HandTransforms"`
}

/* TurnsState
   public bool Enable;
   public TurnType Type = TurnType.Auto;
   public List<string> TurnOrder = new List<string>();
   public bool Reverse;
   public bool SkipEmpty;
   public bool DisableInteractions;
   public bool PassTurns = true;
   public string TurnColor;*/
type Turns struct {
	Enable              bool     `json:"Enable"`
	Type                int      `json:"Type"` //Auto[0] = turn order is based on positioning of hands on around table, Custom[1] = turn order is based on an user color list
	TurnOrder           []string `json:"TurnOrder"`
	Reverse             bool     `json:"Reverse"`
	SkipEmpty           bool     `json:"SkipEmpty"`
	DisableInteractions bool     `json:"DisableInteractions"`
	PassTurns           bool     `json:"PassTurns"`
	TurnColor           string   `json:"TurnColor"`
}

/* TextState
    public string Text;
    public ColourState colorstate;
	public int fontSize = 64;*/
type Text struct {
	Text     string `json:"Text"`
	Color    Color  `json:"colorstate"`
	FontSize int    `json:"fontSize"`
}

/* TabState
   public string title;
   public string body;
   public string color;
   public ColourState visibleColor;
   public int id = -1;*/
type Tab struct {
	Title        string `json:"title"`
	Body         string `json:"body"`
	Color        string `json:"color"`
	VisibleColor Color  `json:"visibleColor"`
	ID           int    `json:"id"`
}

/* SnapPointState
   public VectorState Position; //World position when not attached and local position when attached to an object
   public VectorState Rotation; //Rotate is only set for rotation snap points*/
type SnapPoint struct {
	Position Vector `json:"Position"`
	Rotation Vector `json:"Rotation"`
}

/* DecalState
   public TransformState Transform;
   public CustomDecalState CustomDecal;*/
type Decal struct {
	Transform   Transform   `json:"Transform"`
	CustomDecal CustomDecal `json:"CustomDecal"`
}

/* CustomDecalState
{
    public string Name;
    public string ImageURL;
    public float Size; //Size in inches*/
type CustomDecal struct {
	Name     string  `json:"Name"`
	ImageURL string  `json:"ImageURL"`
	Size     float64 `json:"Size"`
}

/* CustomAssetState
    public string Name;
    [Tag(TagType.URL)]
	public string URL;*/
type CustomAsset struct {
	Name string `json:"Name"`
	URL  string `json:"URL"`
}

/* TransformState
   public float posX;
   public float posY;
   public float posZ;

   public float rotX;
   public float rotY;
   public float rotZ;

   public float scaleX;
   public float scaleY;
   public float scaleZ;*/
type Transform struct {
	PosX   float64 `json:"posX"`
	PosY   float64 `json:"posY"`
	PosZ   float64 `json:"posZ"`
	RotX   float64 `json:"rotX"`
	RotY   float64 `json:"rotY"`
	RotZ   float64 `json:"rotZ"`
	ScaleX float64 `json:"scaleX"`
	ScaleY float64 `json:"scaleY"`
	ScaleZ float64 `json:"scaleZ"`
}

/* ColourState
   public float r, g, b;*/
type Color struct {
	R float64 `json:"r"`
	G float64 `json:"g"`
	B float64 `json:"b"`
	A float64 `json:"a,omitempty"`
}

type MusicPlayer struct {
	RepeatSong        bool        `json:"RepeatSong"`
	PlaylistEntry     int         `json:"PlaylistEntry"`
	CurrentAudioTitle string      `json:"CurrentAudioTitle"`
	CurrentAudioURL   string      `json:"CurrentAudioURL"`
	AudioLibrary      []AudioItem `json:"AudioLibrary"`
}

type AudioItem struct {
	Item1 string `json:"Item1"`
	Item2 string `json:"Item2"`
}

/* CustomDeckState
   public string FaceURL = "";
   public string BackURL = "";
   public int? NumWidth;
   public int? NumHeight;
   public bool BackIsHidden = false; //Back of cards becames the hidden card when in a hand
   public bool UniqueBack = false; //Each back is a unique card just like the front*/
type CustomDeck struct {
	FaceURL      string `json:"FaceURL"`
	BackURL      string `json:"BackURL"`
	NumWidth     int    `json:"NumWidth,omitempty"`
	NumHeight    int    `json:"NumHeight,omitempty"`
	BackIsHidden bool   `json:"BackIsHidden"`
	UniqueBack   bool   `json:"UniqueBack"`
}

/* CustomImageState
   public string ImageURL = "";
   public string ImageSecondaryURL = "";
   public float WidthScale; //Holds the scaled size of the object based on the image dimensions
   public CustomDiceState CustomDice;
   public CustomTokenState CustomToken;
   public CustomJigsawPuzzleState CustomJigsawPuzzle;
   public CustomTileState CustomTile;*/
type CustomImage struct {
	ImageURL           string             `json:"ImageURL"`
	ImageSecondaryURL  string             `json:"ImageSecondaryURL"`
	ImageScalar        int                `json:"ImageScalar"`
	WidthScale         float64            `json:"WidthScale"`
	CustomDice         CustomDice         `json:"CustomDice,omitempty"`
	CustomToken        CustomToken        `json:"CustomToken,omitempty"`
	CustomJigsawPuzzle CustomJigsawPuzzle `json:"CustomJigsawPuzzle,omitempty"`
	CustomTile         CustomTile         `json:"CustomTile,omitempty"`
}

/* CustomAssetbundleState
   public string AssetbundleURL = "";
   public string AssetbundleSecondaryURL = "";
   public int MaterialIndex = 0; //0 = Plastic, 1 = Wood, 2 = Metal, 3 = Cardboard
   public int TypeIndex = 0; //0 = Generic, 1 = Figurine, 2 = Dice, 3 = Coin, 4 = Board, 5 = Chip, 6 = Bag, 7 = Infinite
   public int LoopingEffectIndex = 0;*/
type CustomAssetbundle struct {
	AssetbundleURL          string `json:"AssetbundleURL"`
	AssetbundleSecondaryURL string `json:"AssetbundleSecondaryURL"`
	MaterialIndex           int    `json:"MaterialIndex"` //0 = Plastic, 1 = Wood, 2 = Metal, 3 = Cardboard
	TypeIndex               int    `json:"TypeIndex"`     //0 = Generic, 1 = Figurine, 2 = Dice, 3 = Coin, 4 = Board, 5 = Chip, 6 = Bag, 7 = Infinite
	LoopingEffectIndex      int    `json:"LoopingEffectIndex"`
}

/* CustomDiceState
   public DiceType Type;*/
type CustomDice struct {
	Type int `json:"Type"`
}

/* CustomTokenState
   public float Thickness;
   public float MergeDistancePixels;
   public bool Stackable = false;*/
type CustomToken struct {
	Thickness           float64 `json:"Thickness"`
	MergeDistancePixels float64 `json:"MergeDistancePixels"`
	Stackable           bool    `json:"Stackable"`
}

/* CustomTileState
   public TileType Type; //0 = Box, 1 = Hex, 2 = Circle, 3 = Rounded
   public float Thickness;
   public bool Stackable = false;
   public bool Stretch = false;*/
type CustomTile struct {
	Type      int     `json:"Type"` //0 = Box, 1 = Hex, 2 = Circle, 3 = Rounded
	Thickness float64 `json:"Thickness"`
	Stackable bool    `json:"Stackable"`
	Stretch   bool    `json:"Stretch"`
}

/* CustomJigsawPuzzleState
   public int NumPuzzlePieces = 80;
   public bool ImageOnBoard = true;*/
type CustomJigsawPuzzle struct {
	NumPuzzlePieces int  `json:"NumPuzzlePieces"`
	ImageOnBoard    bool `json:"ImageOnBoard"`
}

/* CustomMeshState
   public string MeshURL = "";
   public string DiffuseURL = "";
   public string NormalURL = "";
   public string ColliderURL = "";
   public bool Convex = true;
   public int MaterialIndex = 0; //0 = Plastic, 1 = Wood, 2 = Metal, 3 = Cardboard
   public int TypeIndex = 0; //0 = Generic, 1 = Figurine, 2 = Dice, 3 = Coin, 4 = Board, 5 = Chip, 6 = Bag, 7 = Infinite
   public CustomShaderState CustomShader; //Used to override the shader
   public bool CastShadows = true;*/
type CustomMesh struct {
	MeshURL       string       `json:"MeshURL"`
	DiffuseURL    string       `json:"DiffuseURL"`
	NormalURL     string       `json:"NormalURL"`
	ColliderURL   string       `json:"ColliderURL"`
	Convex        bool         `json:"Convex"`
	MaterialIndex int          `json:"MaterialIndex"` //0 = Plastic, 1 = Wood, 2 = Metal, 3 = Cardboard
	TypeIndex     int          `json:"TypeIndex"`     //0 = Generic, 1 = Figurine, 2 = Dice, 3 = Coin, 4 = Board, 5 = Chip, 6 = Bag, 7 = Infinite
	CustomShader  CustomShader `json:"CustomShader"`
	CastShadows   bool         `json:"CastShadows"`
}

/* CustomShaderState
   public ColourState SpecularColor = new ColourState(0.9f, 0.9f, 0.9f);
   public float SpecularIntensity = 0.1f;
   public float SpecularSharpness = 3f; //Range: 2 - 8
   public float FresnelStrength = 0.1f; //Range: 0 - 1*/
type CustomShader struct {
	SpecularColor     Color   `json:"SpecularColor"`
	SpecularIntensity float64 `json:"SpecularIntensity"`
	SpecularSharpness float64 `json:"SpecularSharpness"`
	FresnelStrength   float64 `json:"FresnelStrength"`
}

/* FogOfWarSaveState
   public bool HideGmPointer;
   public bool HideObjects;
   public float Height;
   public Dictionary<string, HashSet<int>> RevealedLocations;*/
type FogOfWar struct {
	HideGmPointer     bool          `json:"HideGmPointer"`
	HideObjects       bool          `json:"HideObjects"`
	Height            float64       `json:"Height"`
	RevealedLocations []interface{} `json:"RevealedLocations"`
}

/* FogOfWarRevealerSaveState
   public bool Active;
   public float Range;
   public string Color;*/
type FogOfWarRevealer struct {
	Active bool    `json:"Active"`
	Range  float64 `json:"Range"`
	Color  string  `json:"Color"`
}

/* TabletState
   public string PageURL = "";*/
type Tablet struct {
	PageURL string `json:"PageURL"`
}

/* ClockSaveState
   public ClockScript.ClockState ClockState;
   public int SecondsPassed = 0;
   public bool Paused = false;*/
type Clock struct {
	ClockState    int  `json:"ClockState"`
	SecondsPassed int  `json:"SecondsPassed"`
	Paused        bool `json:"Paused"`
}

/* CounterState
   public int value = 0;*/
type Counter struct {
	Value int `json:"value"`
}

/* Mp3PlayerState
   public string songTitle = "";
   public string genre = "";
   public float volume = 0.5f;
   public bool isPlaying = false;
   public bool loopOne = false;
   public string menuTitle = "GENRES";
   public Menus menu = Menus.GENRES;*/
type Mp3Player struct {
	SongTitle string  `json:"songTitle"`
	Genre     string  `json:"genre"`
	Volume    float64 `json:"volume"`
	IsPlaying bool    `json:"isPlaying"`
	LoopOne   bool    `json:"loopOne"`
	MenuTitle int     `json:"menuTitle"`
	Menu      int     `json:"menu"`
}

/* CalculatorState
   public string value = "";
   public float memory = 0;*/
type Calculator struct {
	Value  int     `json:"value"`
	Memory float64 `json:"memory"`
}

/*VectorLineState
{
    public List<VectorState> points3;
    public ColourState color;
    public float thickness = 0.1f;
    public VectorState rotation;
    public bool? loop;
    public bool? square;*/
type VectorLine struct {
	Points3   []Vector `json:"points3"`
	Color     Color    `json:"color"`
	Thickness float64  `json:"thickness"`
	Rotation  Vector   `json:"rotation"`
	Loop      bool     `json:"loop"`
	Square    bool     `json:"square"`
}

/* JointFixedState : JointState //http://docs.unity3d.com/ScriptReference/FixedJoint.html
   public string ConnectedBodyGUID = ""; //A reference to another rigidbody this joint connects to.
   public bool EnableCollision; //Enable collision between bodies connected with the joint.
   public VectorState Axis = new VectorState(); //The Direction of the axis around which the body is constrained.
   public VectorState Anchor = new VectorState(); //The Position of the anchor around which the joints motion is constrained.
   public VectorState ConnectedAnchor = new VectorState(); //Position of the anchor relative to the connected Rigidbody.
   public float BreakForce; //The force that needs to be applied for this joint to break.
   public float BreakTorgue; //The torque that needs to be applied for this joint to break.*/
type JointFixed struct {
	ConnectedBodyGUID string  `json:"ConnectedBodyGUID"`
	EnableCollision   bool    `json:"EnableCollision"`
	Axis              Vector  `json:"Axis"`
	Anchor            Vector  `json:"Anchor"`
	ConnectedAnchor   Vector  `json:"ConnectedAnchor"`
	BreakForce        float64 `json:"BreakForce"`
	BreakTorgue       float64 `json:"BreakTorgue"`
}

/* JointHingeState : JointState //http://docs.unity3d.com/ScriptReference/HingeJoint.html
    public bool UseLimits;
    public JointLimits Limits; //Limit of angular rotation on the hinge joint. http://docs.unity3d.com/ScriptReference/JointLimits.html
    public bool UseMotor;
    public JointMotor Motor; //The motor will apply a force up to a maximum force to achieve the target velocity in degrees per second. http://docs.unity3d.com/ScriptReference/JointMotor.html
    public bool UseSpring;
    public JointSpring Spring; //The spring attempts to reach a target angle by adding spring and damping forces. http://docs.unity3d.com/ScriptReference/JointSpring.html
	public string ConnectedBodyGUID = ""; //A reference to another rigidbody this joint connects to.
   public bool EnableCollision; //Enable collision between bodies connected with the joint.
   public VectorState Axis = new VectorState(); //The Direction of the axis around which the body is constrained.
   public VectorState Anchor = new VectorState(); //The Position of the anchor around which the joints motion is constrained.
   public VectorState ConnectedAnchor = new VectorState(); //Position of the anchor relative to the connected Rigidbody.
   public float BreakForce; //The force that needs to be applied for this joint to break.
   public float BreakTorgue; //The torque that needs to be applied for this joint to break.*/
type JointHinge struct {
	UseLimits         bool        `json:"UseLimits"`
	UseMotor          bool        `json:"UseMotor"`
	UseSpring         bool        `json:"UseSpring"`
	Spring            JointSpring `json:"Spring"`
	ConnectedBodyGUID string      `json:"ConnectedBodyGUID"`
	EnableCollision   bool        `json:"EnableCollision"`
	Axis              Vector      `json:"Axis"`
	Anchor            Vector      `json:"Anchor"`
	ConnectedAnchor   Vector      `json:"ConnectedAnchor"`
	BreakForce        float64     `json:"BreakForce"`
	BreakTorgue       float64     `json:"BreakTorgue"`
}

/* JointSpringState : JointState //http://docs.unity3d.com/ScriptReference/SpringJoint.html
	public float Damper; //The damper force used to dampen the spring force.
    public float MaxDistance; //The maximum distance between the bodies relative to their initial distance.
    public float MinDistance; //The minimum distance between the bodies relative to their initial distance.
    public float Spring; //The spring force used to keep the two objects together.
   public string ConnectedBodyGUID = ""; //A reference to another rigidbody this joint connects to.
   public bool EnableCollision; //Enable collision between bodies connected with the joint.
   public VectorState Axis = new VectorState(); //The Direction of the axis around which the body is constrained.
   public VectorState Anchor = new VectorState(); //The Position of the anchor around which the joints motion is constrained.
   public VectorState ConnectedAnchor = new VectorState(); //Position of the anchor relative to the connected Rigidbody.
   public float BreakForce; //The force that needs to be applied for this joint to break.
   public float BreakTorgue; //The torque that needs to be applied for this joint to break.*/
type JointSpring struct {
	Damper            float64 `json:"Damper"`
	MaxDistance       float64 `json:"MaxDistance"`
	MinDistance       Vector  `json:"MinDistance"`
	Spring            float64 `json:"Spring"`
	ConnectedBodyGUID string  `json:"ConnectedBodyGUID"`
	EnableCollision   bool    `json:"EnableCollision"`
	Axis              Vector  `json:"Axis"`
	Anchor            Vector  `json:"Anchor"`
	ConnectedAnchor   Vector  `json:"ConnectedAnchor"`
	BreakForce        float64 `json:"BreakForce"`
	BreakTorgue       float64 `json:"BreakTorgue"`
}

/* RigidbodyState
   public float Mass = 1f; //The mass of the object (arbitrary units). You should not make masses more or less than 100 times that of other Rigidbodies.
   public float Drag = 0.1f; //How much air resistance affects the object when moving from forces. 0 means no air resistance, and infinity makes the object stop moving immediately.
   public float AngularDrag = 0.1f; //How much air resistance affects the object when rotating from torque. 0 means no air resistance. Note that you cannot make the object stop rotating just by setting its Angular Drag to infinity.
   public bool UseGravity = true; //If enabled, the object is affected by gravity.*/
type Rigidbody struct {
	Mass        float64 `json:"Mass"`
	Drag        float64 `json:"Drag"`
	AngularDrag float64 `json:"AngularDrag"`
	UseGravity  bool    `json:"UseGravity"`
}

/* public class PhysicsMaterialState
   public float StaticFriction = 0.4f; //The friction used when an object is laying still on a surface. Usually a value from 0 to 1. A value of zero feels like ice, a value of 1 will make it very hard to get the object moving.
   public float DynamicFriction = 0.4f; //The friction used when already moving. Usually a value from 0 to 1. A value of zero feels like ice, a value of 1 will make it come to rest very quickly unless a lot of force or gravity pushes the object.
   public float Bounciness = 0f; //How bouncy is the surface? A value of 0 will not bounce. A value of 1 will bounce without any loss of energy.
   public PhysicMaterialCombine FrictionCombine = PhysicMaterialCombine.Average; //How the friction of two colliding objects is combined. 0 = Average, 1 = Minimum, 2 = Maximum, 3 = Multiply.
   public PhysicMaterialCombine BounceCombine = PhysicMaterialCombine.Average; //How the bounciness of two colliding objects is combined. 0 = Average, 1 = Minimum, 2 = Maximum, 3 = Multiply.*/
type PhysicsMaterial struct {
	StaticFriction  float64 `json:"StaticFriction"`
	DynamicFriction float64 `json:"DynamicFriction"`
	Bounciness      float64 `json:"Bounciness"`
	FrictionCombine int     `json:"FrictionCombine"` // 0 = Average, 1 = Minimum, 2 = Maximum, 3 = Multiply.
	BounceCombine   int     `json:"BounceCombine"`   // 0 = Average, 1 = Minimum, 2 = Maximum, 3 = Multiply.
}

/* HandTransformState
   public string Color;
   public TransformState Transform;*/
type HandTransform struct {
	Color     string    `json:"Color"`
	Transform Transform `json:"Transform"`
}

type DecalPallet struct {
	Name     string  `json:"Name"`
	ImageURL string  `json:"ImageURL"`
	Size     float64 `json:"Size"`
}

type CustomPDF struct {
	PDFURL        string `json:"PDFUrl"`
	PDFPassword   string `json:"PDFPassword"`
	PDFPage       int    `json:"PDFPage"`
	PDFPageOffset int    `json:"PDFPageOffset"`
}

/*CameraState
  public VectorState Position;
  public VectorState Rotation;
  public float Distance;
  public bool Zoomed = false;*/
type CameraState struct {
	Position         Vector  `json:"Position"`
	Rotation         Vector  `json:"Rotation"`
	Distance         float64 `json:"Distance"`
	Zoomed           bool    `json:"Zoomed"`
	AbsolutePosition Vector  `json:"AbsolutePosition"`
}

type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
