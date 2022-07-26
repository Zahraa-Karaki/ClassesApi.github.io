PGDMP     +                    z           stc    14.4    14.4                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    16448    stc    DATABASE     g   CREATE DATABASE stc WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_United States.1252';
    DROP DATABASE stc;
                postgres    false            �            1259    16449    classes    TABLE     �   CREATE TABLE public.classes (
    id integer NOT NULL,
    name character varying(64) NOT NULL,
    language character varying(64)
);
    DROP TABLE public.classes;
       public         heap    postgres    false            �            1259    16474    class_id_seq    SEQUENCE     �   ALTER TABLE public.classes ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.class_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    209            �            1259    16464    students    TABLE     �   CREATE TABLE public.students (
    id integer NOT NULL,
    firstname character varying(128) NOT NULL,
    lastname character varying(128) NOT NULL,
    classid integer
);
    DROP TABLE public.students;
       public         heap    postgres    false            �            1259    16475    student_id_seq    SEQUENCE     �   ALTER TABLE public.students ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.student_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    211            �            1259    16454    teachers    TABLE     �   CREATE TABLE public.teachers (
    id integer NOT NULL,
    firstname character varying(128) NOT NULL,
    lastname character varying(128) NOT NULL,
    classid integer
);
    DROP TABLE public.teachers;
       public         heap    postgres    false            �            1259    16476    teacher_id_seq    SEQUENCE     �   ALTER TABLE public.teachers ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.teacher_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    210            �          0    16449    classes 
   TABLE DATA           5   COPY public.classes (id, name, language) FROM stdin;
    public          postgres    false    209   g       �          0    16464    students 
   TABLE DATA           D   COPY public.students (id, firstname, lastname, classid) FROM stdin;
    public          postgres    false    211   �       �          0    16454    teachers 
   TABLE DATA           D   COPY public.teachers (id, firstname, lastname, classid) FROM stdin;
    public          postgres    false    210   �                  0    0    class_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.class_id_seq', 3, true);
          public          postgres    false    212                       0    0    student_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.student_id_seq', 4, true);
          public          postgres    false    213                       0    0    teacher_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.teacher_id_seq', 5, true);
          public          postgres    false    214            g           2606    16453    classes class_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.classes
    ADD CONSTRAINT class_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.classes DROP CONSTRAINT class_pkey;
       public            postgres    false    209            k           2606    16468    students student_pkey 
   CONSTRAINT     S   ALTER TABLE ONLY public.students
    ADD CONSTRAINT student_pkey PRIMARY KEY (id);
 ?   ALTER TABLE ONLY public.students DROP CONSTRAINT student_pkey;
       public            postgres    false    211            i           2606    16458    teachers teacher_pkey 
   CONSTRAINT     S   ALTER TABLE ONLY public.teachers
    ADD CONSTRAINT teacher_pkey PRIMARY KEY (id);
 ?   ALTER TABLE ONLY public.teachers DROP CONSTRAINT teacher_pkey;
       public            postgres    false    210            m           2606    16469    students FK_class_student    FK CONSTRAINT     �   ALTER TABLE ONLY public.students
    ADD CONSTRAINT "FK_class_student" FOREIGN KEY (classid) REFERENCES public.classes(id) ON UPDATE SET NULL ON DELETE SET NULL;
 E   ALTER TABLE ONLY public.students DROP CONSTRAINT "FK_class_student";
       public          postgres    false    209    211    3175            l           2606    16459    teachers FK_class_teacher    FK CONSTRAINT     �   ALTER TABLE ONLY public.teachers
    ADD CONSTRAINT "FK_class_teacher" FOREIGN KEY (classid) REFERENCES public.classes(id) ON UPDATE SET NULL ON DELETE SET NULL;
 E   ALTER TABLE ONLY public.teachers DROP CONSTRAINT "FK_class_teacher";
       public          postgres    false    210    3175    209            �   .   x�3�L/JLI5�L�K��,��2�s������\A� ��      �   9   x�ƻ�0@��1L$�'	(ٿq����bf,ٸ;�����[^��������]�g      �   8   x�3�LLJcC.#���|���LNc.���bΔ� ۔��������ӈ+F��� h��     