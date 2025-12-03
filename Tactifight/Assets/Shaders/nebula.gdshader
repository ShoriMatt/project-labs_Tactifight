shader_type canvas_item;

uniform float intensity = 0.6;
uniform float scale     = 1.2;
uniform float flow      = 0.20;
uniform float contrast  = 1.2;
uniform vec3  colA = vec3(0.20, 0.06, 0.40);
uniform vec3  colB = vec3(0.75, 0.35, 0.80);

float hash(vec2 p){
    p = vec2(dot(p, vec2(127.1,311.7)), dot(p, vec2(269.5,183.3)));
    return fract(sin(p.x+p.y)*43758.5453);
}
float noise(vec2 p){
    vec2 i = floor(p);
    vec2 f = fract(p);
    float a = hash(i);
    float b = hash(i+vec2(1,0));
    float c = hash(i+vec2(0,1));
    float d = hash(i+vec2(1,1));
    vec2 u = f*f*(3.0-2.0*f);
    return mix(mix(a,b,u.x), mix(c,d,u.x), u.y);
}
float fbm(vec2 p){
    float v=0.0, a=0.5;
    for (int i=0;i<3;i++){ v += a*noise(p); p*=2.0; a*=0.5; }
    return v;
}

void fragment(){
    vec2 uv = (UV - 0.5) * scale + 0.5;
    uv += vec2(TIME * 0.03 * flow, -TIME * 0.02 * flow);

    float n1 = fbm(uv * 2.0);
    float n2 = fbm((uv + 10.0) * 1.5);
    float m = pow(smoothstep(0.25, 0.85, (n1*0.6 + n2*0.4)), contrast);

    vec3 neb = mix(colA, colB, m) * intensity;

    float vignette = smoothstep(0.95, 0.35, length(UV - 0.5));
    vec3 bg = vec3(0.010, 0.012, 0.020);
    vec3 color = mix(bg, neb, vignette);
    COLOR = vec4(color, 1.0);
}
