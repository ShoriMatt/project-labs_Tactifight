shader_type canvas_item;

uniform float density = 0.92;
uniform float twinkle = 0.25;
uniform vec2  drift   = vec2(0.03, 0.0);

float hash(vec2 p){
    p = vec2(dot(p, vec2(127.1,311.7)), dot(p, vec2(269.5,183.3)));
    return fract(sin(p.x+p.y)*43758.5453);
}

void fragment(){
    vec2 uv = UV + drift * TIME;
    vec2 gv = fract(uv * 170.0) - 0.5;
    vec2 id = floor(uv * 170.0);

    float rnd = hash(id);
    vec2  starp = vec2(fract(rnd*97.0), fract(rnd*57.0)) - 0.5;

    float d = length(gv - starp);
    float star = smoothstep(0.08, 0.0, d);
    float glow = smoothstep(0.28, 0.0, d);

    float tw = 1.0 + twinkle * sin(TIME*6.0 + rnd*20.0);
    float mask = step(1.0 - density, rnd);

    vec3 bg = mix(vec3(0.01,0.01,0.015), vec3(0.04,0.02,0.06), uv.y);
    vec3 col = bg + (star*1.4 + glow*0.25) * tw * mask;
    COLOR = vec4(col, 1.0);
}
