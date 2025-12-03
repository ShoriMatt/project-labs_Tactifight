shader_type canvas_item;

uniform vec2  center = vec2(0.75, 0.25);
uniform vec2  dir    = vec2(-1.0, -0.2);
uniform float tail   = 0.35;
uniform float core   = 0.015;
uniform vec3  color  = vec3(1.0, 0.95, 0.75);

void fragment(){
    vec2  uv = UV;
    vec2  p  = uv - center;

    vec2  n  = normalize(dir);
    float along = dot(p, n);
    float across = dot(p, vec2(-n.y,n.x));

    float coreMask = smoothstep(core, 0.0, length(p));
    float trail = smoothstep(0.0, tail, -along) * smoothstep(0.20, 0.0, abs(across));
    trail *= smoothstep(0.0, 0.6, -along);
    float glow = smoothstep(0.30, 0.0, length(p)) * 0.5;

    vec3 col = color * (coreMask*1.0 + trail*1.2 + glow*0.7);
    COLOR = vec4(col, max(coreMask, max(trail, glow)));
}
