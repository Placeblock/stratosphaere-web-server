INSERT INTO articles (title, description, content, author) VALUES ("Der erste Blogbeitrag", "Schau dir diesen wundervollen ersten Blogbeitrag an :)", '<h1>Der erste Blogbeitrag!</h1><p><br></p><p>Wenn du das hier siehst, dann weist du schonmal, dass der Blog funktioniert!</p><p>Bitte bedenke, dass damit du diesen Blogbeitrag siehst viel in deinem Computer passiert.</p><p>Was passiert musste ich ihm erst einmal beibringen, darin stecken <strong style="color: rgb(230, 0, 0);">viele Stunden Arbeit</strong>!</p><p><br></p><p><strong><u>Was musste gemacht werden?</u></strong></p><ul><li>Das Login System damit auch nur Authentifizierte Benutzer Blogbeiträge bearbeiten können</li><li>Eine REST-API mit der dein "Mobiles Endgerät" kommunizieren kann</li><li>Die Datenbank Struktur, damit auch alles schön abgespeichert wird!</li><li>Das Frontend der Webseite, also das was du siehst</li><li>Das Laden der Daten ins Frontend</li><li>Und noch viiieeeles mehr</li></ul><p><br></p><p><span class="ql-size-large">Hier sind ein paar dinge die man in Blogposts verwenden kann:</span></p><p><br></p><p><strong><u>Code Highlighting:</u></strong></p><pre class="ql-syntax" spellcheck="false"><span class="hljs-function"><span class="hljs-keyword">func</span> <span class="hljs-params">(a Auth)</span></span> Check() (<span class="hljs-type">bool</span>, <span class="hljs-type">error</span>) {
&nbsp; &nbsp; <span class="hljs-keyword">var</span> auth Auth
&nbsp; &nbsp; err := db.Select(<span class="hljs-string">"id, password"</span>).Where(Auth{Username: a.Username}).First(&amp;auth).Error
&nbsp; &nbsp; <span class="hljs-keyword">if</span> err != <span class="hljs-literal">nil</span> &amp;&amp; err != gorm.ErrRecordNotFound {
&nbsp; &nbsp; &nbsp; &nbsp; <span class="hljs-keyword">return</span> <span class="hljs-literal">false</span>, err
&nbsp; &nbsp; }
&nbsp; &nbsp; a.ID = <span class="hljs-type">uint16</span>(auth.ID)
&nbsp; &nbsp; <span class="hljs-keyword">return</span> util.CompareHash(auth.Password, []<span class="hljs-type">byte</span>(a.Password)), <span class="hljs-literal">nil</span>
}
</pre><p><br></p><p><strong><u>Mathematische Ausdrücke:</u></strong></p><p><span class="ql-formula" data-value="% \f is defined as #1f(#2) using the macro \f\relax{x} = \int_{-\infty}^\infty     \f\hat\xi\,e^{2 \pi i \xi x}     \,d\xi">﻿<span contenteditable="false"><span class="katex"><span class="katex-mathml"><math xmlns="http://www.w3.org/1998/Math/MathML"><semantics><mrow></mrow><annotation encoding="application/x-tex">% \f is defined as #1f(#2) using the macro \f\relax{x} = \int_{-\infty}^\infty     \f\hat\xi\,e^{2 \pi i \xi x}     \,d\xi</annotation></semantics></math></span><span class="katex-html" aria-hidden="true"></span></span></span>﻿</span> <span class="ql-formula" data-value="% \f is defined as #1f(#2) using the macro \f\relax{x} = \int_{-\infty}^\infty\f\hat\xi\,e^{2 \pi i \xi x}\,d\xi">﻿<span contenteditable="false"><span class="katex"><span class="katex-mathml"><math xmlns="http://www.w3.org/1998/Math/MathML"><semantics><mrow></mrow><annotation encoding="application/x-tex">% \f is defined as #1f(#2) using the macro \f\relax{x} = \int_{-\infty}^\infty\f\hat\xi\,e^{2 \pi i \xi x}\,d\xi</annotation></semantics></math></span><span class="katex-html" aria-hidden="true"></span></span></span>﻿</span> <span class="ql-formula" data-value="\displaystyle \left( \sum_{k=1}^n a_k b_k \right)^2 \leq \left( \sum_{k=1}^n a_k^2 \right) \left( \sum_{k=1}^n b_k^2 \right)">﻿<span contenteditable="false"><span class="katex"><span class="katex-mathml"><math xmlns="http://www.w3.org/1998/Math/MathML">
<semantics><mrow><mstyle scriptlevel="0" displaystyle="true"><msup><mrow><mo fence="true">(</mo><munderover><mo>∑</mo><mrow><mi>k</mi><mo>=</mo><mn>1</mn></mrow><mi>n</mi></munderover><msub><mi>a</mi><mi>k</mi></msub><msub><mi>b</mi><mi>k</mi></msub><mo fence="true">)</mo></mrow><mn>2</mn></msup><mo>≤</mo><mrow><mo fence="true">(</mo><munderover><mo>∑</mo><mrow><mi>k</mi><mo>=</mo><mn>1</mn></mrow><mi>n</mi></munderover><msubsup><mi>a</mi><mi>k</mi><mn>2</mn></msubsup><mo fence="true">)</mo></mrow><mrow><mo fence="true">(</mo><munderover><mo>∑</mo><mrow><mi>k</mi><mo>=</mo><mn>1</mn></mrow><mi>n</mi></munderover><msubsup><mi>b</mi><mi>k</mi><mn>2</mn></msubsup><mo fence="true">)</mo></mrow></mstyle></mrow><annotation encoding="application/x-tex">\displaystyle \left( \sum_{k=1}^n a_k b_k \right)^2 \leq \left( \sum_{k=1}^n a_k^2 \right) \left( \sum_{k=1}^n b_k^2 \right)</annotation></semantics></math></span><span class="katex-html" aria-hidden="true"><span class="base"><span class="strut" style="height: 3.2561em; vertical-align: -1.3021em;"></span><span class="minner"><span class="minner"><span class="mopen delimcenter" style="top: 0em;"><span class="delimsizing size4">(</span></span><span class="mop op-limits"><span class="vlist-t vlist-t2"><span class="vlist-r"><span class="vlist" style="height: 1.6514em;"><span class="" style="top: -1.8479em; margin-left: 0em;"><span class="pstrut" style="height: 3.05em;"></span><span class="sizing reset-size6 size3 mtight"><span class="mord mtight"><span style="margin-right: 0.0315em;" class="mord mathnormal mtight">k</span><span class="mrel mtight">=</span><span class="mord mtight">1</span></span></span></span><span class="" style="top: -3.05em;"><span class="pstrut" style="height: 3.05em;"></span><span class=""><span class="mop op-symbol large-op">∑</span></span></span><span class="" style="top: -4.3em; margin-left: 0em;"><span class="pstrut" style="height: 3.05em;"></span><span class="sizing reset-size6 size3 mtight">
<span class="mord mathnormal mtight">n</span></span></span></span><span class="vlist-s">​</span></span><span class="vlist-r"><span class="vlist" style="height: 1.3021em;"><span class=""></span></span></span></span></span><span class="mspace" style="margin-right: 0.1667em;"></span><span class="mord"><span class="mord mathnormal">a</span><span class="msupsub"><span class="vlist-t vlist-t2"><span class="vlist-r"><span class="vlist" style="height: 0.3361em;"><span class="" style="top: -2.55em; margin-left: 0em; margin-right: 0.05em;"><span class="pstrut" style="height: 2.7em;"></span><span class="sizing reset-size6 size3 mtight"><span style="margin-right: 0.0315em;" class="mord mathnormal mtight">k</span></span></span></span><span class="vlist-s">​</span></span><span class="vlist-r"><span class="vlist" style="height: 0.15em;"><span class=""></span></span></span></span></span></span><span class="mord"><span class="mord mathnormal">b</span><span class="msupsub"><span class="vlist-t vlist-t2"><span class="vlist-r"><span class="vlist" style="height: 0.3361em;"><span class="" style="top: -2.55em; margin-left: 0em; margin-right: 0.05em;"><span class="pstrut" style="height: 2.7em;"></span><span class="sizing reset-size6 size3 mtight"><span style="margin-right: 0.0315em;" class="mord mathnormal mtight">k</span></span></span></span><span class="vlist-s">​</span></span><span class="vlist-r"><span class="vlist" style="height: 0.15em;"><span class=""></span></span></span></span></span></span><span class="mclose delimcenter" style="top: 0em;"><span class="delimsizing size4">)</span></span></span><span class="msupsub"><span class="vlist-t"><span class="vlist-r"><span class="vlist" style="height: 1.954em;"><span class="" style="top: -4.2029em; margin-right: 0.05em;"><span class="pstrut" style="height: 2.7em;"></span><span class="sizing reset-size6 size3 mtight"><span class="mord mtight">2</span></span></span></span></span></span></span></span><span class="mspace" style="margin-right: 0.2778em;"></span><span class="mrel">≤</span><span class="mspace" style="margin-right: 0.2778em;"></span></span><span class="base"><span class="strut" style="height: 3.0521em; vertical-align: -1.3021em;"></span><span class="minner"><span class="mopen delimcenter" style="top: 0em;"><span class="delimsizing size4">(</span></span><span class="mop op-limits"><span class="vlist-t vlist-t2"><span class="vlist-r"><span class="vlist" style="height: 1.6514em;"><span class="" style="top: -1.8479em; margin-left: 0em;"><span class="pstrut" style="height: 3.05em;"></span><span class="sizing reset-size6 size3 mtight"><span class="mord mtight"><span style="margin-right: 0.0315em;" class="mord mathnormal mtight">k</span><span class="mrel mtight">=</span><span class="mord mtight">1</span></span></span></span><span class="" style="top: -3.05em;"><span class="pstrut" style="height: 3.05em;"></span><span class=""><span class="mop op-symbol large-op">∑</span></span></span><span class="" style="top: -4.3em; margin-left: 0em;"><span class="pstrut" style="height: 3.05em;"></span><span class="sizing reset-size6 size3 mtight"><span class="mord mathnormal mtight">n</span>
</span></span></span><span class="vlist-s">​</span></span><span class="vlist-r"><span class="vlist" style="height: 1.3021em;"><span class=""></span></span></span></span></span><span class="mspace" style="margin-right: 0.1667em;"></span><span class="mord"><span class="mord mathnormal">a</span><span class="msupsub"><span class="vlist-t vlist-t2"><span class="vlist-r"><span class="vlist" style="height: 0.8641em;"><span class="" style="top: -2.453em; margin-left: 0em; margin-right: 0.05em;"><span class="pstrut" style="height: 2.7em;"></span><span class="sizing reset-size6 size3 mtight"><span style="margin-right: 0.0315em;" class="mord mathnormal mtight">k</span></span></span><span class="" style="top: -3.113em; margin-right: 0.05em;"><span class="pstrut" style="height: 2.7em;"></span><span class="sizing reset-size6 size3 mtight"><span class="mord mtight">2</span></span></span></span><span class="vlist-s">​</span></span><span class="vlist-r"><span class="vlist" style="height: 0.247em;"><span class=""></span></span></span></span></span></span><span class="mclose delimcenter" style="top: 0em;"><span class="delimsizing size4">)</span></span></span><span class="mspace" style="margin-right: 0.1667em;"></span><span class="minner"><span class="mopen delimcenter" style="top: 0em;"><span class="delimsizing size4">(</span></span><span class="mop op-limits"><span class="vlist-t vlist-t2"><span class="vlist-r"><span class="vlist" style="height: 1.6514em;"><span class="" style="top: -1.8479em; margin-left: 0em;"><span class="pstrut" style="height: 3.05em;"></span><span class="sizing reset-size6 size3 mtight"><span class="mord mtight"><span style="margin-right: 0.0315em;" class="mord mathnormal mtight">k</span><span class="mrel mtight">=</span><span class="mord mtight">1</span></span></span></span><span class="" style="top: -3.05em;"><span class="pstrut" style="height: 3.05em;"></span><span class=""><span class="mop op-symbol large-op">∑</span></span></span><span class="" style="top: -4.3em; margin-left: 0em;"><span class="pstrut" style="height: 3.05em;"></span><span class="sizing reset-size6 size3 mtight"><span class="mord mathnormal mtight">n</span></span></span></span><span class="vlist-s">​</span></span><span class="vlist-r"><span class="vlist" style="height: 1.3021em;"><span class=""></span></span></span></span></span><span class="mspace" style="margin-right: 0.1667em;"></span><span class="mord"><span class="mord mathnormal">b</span><span class="msupsub"><span class="vlist-t vlist-t2"><span class="vlist-r"><span class="vlist" style="height: 0.8641em;"><span class="" style="top: -2.453em; margin-left: 0em; margin-right: 0.05em;"><span class="pstrut" style="height: 2.7em;"></span><span class="sizing reset-size6 size3 mtight"><span style="margin-right: 0.0315em;" class="mord mathnormal mtight">k</span></span></span><span class="" style="top: -3.113em; margin-right: 0.05em;"><span class="pstrut" style="height: 2.7em;"></span><span class="sizing reset-size6 size3 mtight"><span class="mord mtight">2</span></span></span></span><span class="vlist-s">​</span></span><span class="vlist-r"><span class="vlist" style="height: 0.247em;"><span class=""></span></span></span></span></span></span><span class="mclose delimcenter" style="top: 0em;"><span class="delimsizing size4">)</span></span></span></span></span></span></span>﻿</span></p><p><br></p><p><strong><u>Bilder:</u></strong></p><p><img src="file:///C:/Users/xilef/AppData/Roaming/Gajim/downloads/Projekt%20Solis-Dateien/abstract-digital-futuristic-eye-picture-id1322220448.jpeg"></p><p><br></p><p><strong><u>Videos:</u></strong></p><iframe class="ql-video" allowfullscreen="true" src="./Projekt Solis_files/flower_002.html" frameborder="0"></iframe><p><br></p><p>(Irgendein random Video was ich gefunden habe ^^)</p><p><br></p><p><span style="color: rgb(255, 153, 0); background-color: rgb(0, 0, 0);">Und </span>noch <span style="background-color: rgb(178, 107, 0); color: rgb(255, 255, 255);">vieles </span><span style="background-color: rgb(102, 185, 102); color: rgb(230, 0, 0);">mehr </span>;)</p>', "Felix Weglehner");